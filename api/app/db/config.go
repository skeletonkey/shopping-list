// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
package db

import "github.com/skeletonkey/lib-core-go/config"

var cfg *db

func getConfig() *db {
	config.LoadConfig("db", &cfg)
	return cfg
}
