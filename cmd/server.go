package cmd

import (
	"github.com/spf13/cobra"
)

var server = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "Runs the server",
	Long:    "Starts the Pass Around Novel API server",
	Run: func(cmd *cobra.Command, args []string) {
		l.Info("Starting server...")
		l.Debug("Debug message")
		l.Info("Info message")
		l.Warn("Warn message")
		l.Error("Error message")
		l.Info("Multi\nLine\nMessage")
	},
}

func init() {
	root.AddCommand(server)
}
