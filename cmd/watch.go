package cmd

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type configVar struct {
	key     string
	ptr     interface{}
	decoder func(interface{}, interface{})
}

var (
	configMap       []configVar
	hasConfigLoaded bool = false
	// ConfigUpdated contains handlers to call once the configuration file changes or is loaded
	ConfigUpdated []func()
)

func getConfig(key string, val interface{}, decoder func(interface{}, interface{})) {
	if watchConfig || !hasConfigLoaded {
		configMap = append(configMap, configVar{
			key:     key,
			ptr:     val,
			decoder: decoder,
		})
	}
	if hasConfigLoaded {
		decoder(viper.Get(key), val)
	}
}

func updateConfig(e fsnotify.Event) {
	for _, c := range configMap {
		c.decoder(viper.Get(c.key), c.ptr)
	}
	fireConfigUpdated()
	l.Info("Configuration reloaded.")
}

func configLoaded() {
	for _, c := range configMap {
		c.decoder(viper.Get(c.key), c.ptr)
	}
	if !watchConfig {
		configMap = nil
	}
	fireConfigUpdated()
}

func fireConfigUpdated() {
	for _, f := range ConfigUpdated {
		f()
	}
}
