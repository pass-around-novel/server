package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	root = &cobra.Command{
		Short: "Pass Around Novel server",
		Long:  "Runs the server layer for the Pass Around Novel app, which allows users to write stories with their friends",
	}
	cfgFile     string
	watchConfig bool
)

// Execute runs the command-line flag parser and starts the program
func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	root.Use = os.Args[0]
	cobra.OnInitialize(initConfig)
	root.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file to use")
	root.PersistentFlags().BoolVar(&watchConfig, "watch-config", false, "automatically reload configuration when the file changes")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("pass-around-novel.conf")
		viper.SetConfigType("ini")
		viper.AddConfigPath("/etc/")
		viper.AddConfigPath("$HOME/.config")
		viper.AddConfigPath(".")
	}
	viper.SetEnvPrefix("PASS_AROUND_NOVEL")
	viper.AutomaticEnv()
	if watchConfig {
		viper.WatchConfig()
		viper.OnConfigChange(updateConfig)
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Fprintln(os.Stderr, "WARN: No config file found.  Loading defaults.")
		} else {
			fmt.Fprintf(os.Stderr, "Unable to read config file: %s\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Read configuration from %s\n", viper.ConfigFileUsed())
	}
}
