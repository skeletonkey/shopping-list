// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
package server

import "github.com/skeletonkey/lib-core-go/config"

var cfg *server

func getConfig() *server {
	config.LoadConfig("server", &cfg)
	return cfg
}
