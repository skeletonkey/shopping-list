package db

import (
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

func (d *db) Initialize() {
	log := logger.Get()
	log.Trace().Msg("Configuration updated, attempting to create/update families")
	for _, f := range d.Family {
		family := Family{
			Name:        f.Name,
			DisplayName: f.DisplayName,
		}
		existingFamily, err := getFamilyByName(f.Name)
		if err != nil {
			log.Err(err).Str("name", f.Name).Msg("getting family by name")
		}
		if existingFamily.ID != 0 {
			if family.DisplayName != "" && family.DisplayName != existingFamily.DisplayName {
				existingFamily.DisplayName = family.DisplayName
				if err := existingFamily.update(); err != nil {
					log.Error().Err(err).Msg("updating family")
					continue
				}
				log.Info().Int("id", family.ID).Str("display name", existingFamily.DisplayName).Msg("update family")
			}
			continue
		}
		if err := family.create(); err != nil {
			log.Error().Err(err).Str("name", f.Name).Str("displayName", f.DisplayName).Msg("create attempt failed")
			continue
		}
		log.Info().Str("name", f.Name).Str("displayName", f.DisplayName).Msg("created family")
	}
	log.Trace().Msg("Family create/update complete")
}
