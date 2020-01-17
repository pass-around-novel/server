package cmd

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type configVar struct {
	key     string
	ptr     interface{}
	decoder func(interface{}, interface{})
}

var configMap []configVar

func getConfig(key string, val interface{}, decoder func(interface{}, interface{})) {
	if watchConfig {
		configMap = append(configMap, configVar{
			key:     key,
			ptr:     val,
			decoder: decoder,
		})
	}
	decoder(viper.Get(key), val)
}

func updateConfig(e fsnotify.Event) {
	for _, c := range configMap {
		c.decoder(viper.Get(c.key), c.ptr)
	}
	fmt.Println("Configuration reloaded.")
}
