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
	// Verbose is the number of verbosity levels requested by the user
	Verbose int
	// Quiet is the number of quiet levels requested by the user
	Quiet    int
	lateInit []func()
)

// Execute runs the command-line flag parser and starts the program
func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// AddCommand adds a subcommand to Cobra
func AddCommand(cmd *cobra.Command) {
	root.AddCommand(cmd)
}

func init() {
	root.Use = os.Args[0]
	cobra.OnInitialize(initConfig)
	root.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file to use")
	root.PersistentFlags().BoolVar(&watchConfig, "watch-config", false, "automatically reload configuration when the file changes")
	root.PersistentFlags().CountVarP(&Verbose, "verbose", "v", "Make the output more verbose")
	root.PersistentFlags().CountVarP(&Quiet, "quiet", "q", "Make the output less verbose")
}

func initConfig() {
	fireConfigUpdated()
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("pass-around-novel.conf")
		viper.SetConfigType("toml")
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
			configLoaded()
			l.Warn("No config file found.  Loading defaults.")
		} else {
			fmt.Printf("Unable to read config file: %s\n", err)
			os.Exit(1)
		}
	} else {
		configLoaded()
		l.Infof("Read configuration from %s", viper.ConfigFileUsed())
	}
}
