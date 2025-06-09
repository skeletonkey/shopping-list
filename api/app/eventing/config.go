// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
package eventing

import "github.com/skeletonkey/lib-core-go/config"

var cfg *eventing

func getConfig() *eventing {
	config.LoadConfig("eventing", &cfg)
	return cfg
}
