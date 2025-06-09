package eventing

import (
	"os"
	"time"
)

const touchFileSuffix = ".touch"

type cacheEntries map[cache]bool

type events struct {
	cache        cacheEntries
	cacheUpdated bool
}

type cache struct {
	listId int
	itemId int
}

type eventing struct {
	File        string        `json:"file_name"`
	Compression compression   `json:"compression"`
	Cache       cacheSettings `json:"cache"`
}

type compression struct {
	Time              int `json:"time"`
	TimeInterval      int `json:"time_interval"`
	Size              int `json:"size"`
	SizeCheckInterval int `json:"size_check_interval"`
}

type cacheSettings struct {
	TimeBetweenWrite      int `json:"time_between_writes"`
	LockAlertThreshold    int `json:"data_file_lock_alert_threshold"`
	LockOverrideThreshold int `json:"data_file_lock_override_threshold"`
}

// return the touch file name to determine locking
func (e eventing) touchFile() string {
	return e.File + touchFileSuffix
}

// returns the age of the touch file in seconds
// an age of 0 is not support as 0 indicates that the touch file doesn't exists
func (e eventing) touchFileAge() (int, error) {
	fileInfo, err := os.Stat(e.touchFile())
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}

	return int(time.Since(fileInfo.ModTime())), nil
}
