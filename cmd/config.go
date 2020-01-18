package cmd

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("db.dgraph.host", "127.0.0.1")
	viper.SetDefault("db.dgraph.port", 9080)
	viper.SetDefault("log.console.timeFormat", "15:04:05")
	viper.SetDefault("log.console.primaryFormats", []string{
		"[%s] [DBG ] %s\n",
		"[%s] [INFO] %s\n",
		"[%s] [WARN] %s\n",
		"[%s] [ERR ] %s\n",
	})
	viper.SetDefault("log.console.secondaryFormats", []string{
		"           [DBG ] %[2]s\n",
		"           [INFO] %[2]s\n",
		"           [WARN] %[2]s\n",
		"           [ERR ] %[2]s\n",
	})
	viper.SetDefault("log.console.defaultLevel", 1)
	viper.SetDefault("log.console.levelOverrides", map[string]int{})
}
