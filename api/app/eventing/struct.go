package eventing

type eventing struct {
	File string `json:"file_name"`
	Compression compression `json:"compression"`
	Cache cache `json:"cache"`
}

type compression struct {
	Time int `json:"time"`
	TimeInterval int `json:"time_interval"`
	Size int `json:"size"`
	SizeCheckInterval int `json:"size_check_interval"`
}

type cache struct {
	TimeBetweenWrite int `json:"time_between_writes"`
	LockAlertThreshold int `json:"data_file_lock_alert"`
	LockOverrideThreshold int `json:"data_file_lock_override"`
}