package db

import (
	"context"

	"github.com/skeletonkey/lib-core-go/logger"
)

type db struct {
	DbFile string         `json:"location"`
	Family []FamilyConfig `json:"family"`
}

// FamilyConfig represents a family configuration from the config file
type FamilyConfig struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

func (dbCfg *db) Initialize() {
	log := logger.Get()
	log.Trace().Msg("Configuration updated, attempting to create/update families")

	// This will work for a couple of families; if there are many families this may time out.
	queryCtx, cancel := context.WithTimeout(context.Background(), readTimeout)
	defer cancel()

	type familyInfo struct {
		Family
		Used bool
	}

	allDBFamilies, err := GetAllFamilies(queryCtx)
	dbFamilyFound := make(map[string]*familyInfo, len(allDBFamilies))
	for _, f := range allDBFamilies {
		dbFamilyFound[f.Name] = &familyInfo{Family: f, Used: false}
	}

	if err != nil {
		log.Err(err).Msg("getting all families")
		return
	}

familyFromConfig:
	for _, f := range dbCfg.Family {
		cfgFamily := Family{
			Name:        f.Name,
			DisplayName: f.DisplayName,
		}

		existingFamilyInfo, found := dbFamilyFound[cfgFamily.Name]
		if found {
			existingFamilyInfo.Used = true
			if cfgFamily.DisplayName != "" && cfgFamily.DisplayName != existingFamilyInfo.DisplayName {
				existingFamilyInfo.DisplayName = cfgFamily.DisplayName
				if err := existingFamilyInfo.update(queryCtx); err != nil {
					log.Error().Err(err).Msg("updating family")
					continue familyFromConfig
				}
				log.Info().Int("id", cfgFamily.ID).Str("display name", existingFamilyInfo.DisplayName).Msg("update family")
			}
			continue familyFromConfig
		}
		if err := cfgFamily.create(queryCtx); err != nil {
			log.Error().Err(err).Str("name", f.Name).Str("displayName", f.DisplayName).Msg("create attempt failed")
			continue familyFromConfig
		}
		log.Info().Str("name", f.Name).Str("displayName", f.DisplayName).Msg("created family")
	}
	for _, f := range dbFamilyFound {
		if !f.Used {
			if err := f.Family.removeFamily(queryCtx); err != nil {
				log.Error().Err(err).Int("id", f.ID).Str("name", f.Name).Msg("failed to delete unused family")
			} else {
				log.Info().Int("id", f.ID).Str("name", f.Name).Msg("deleted unused family")
			}
		}
	}

	log.Trace().Msg("Family create/delete/update complete")
}
