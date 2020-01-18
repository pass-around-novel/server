package server

import (
	"os"

	"../cmd"
	"../db/graph"
	"github.com/spf13/cobra"
)

var c = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "Runs the server",
	Long:    "Starts the Pass Around Novel API server",
	Run: func(cmd *cobra.Command, args []string) {
		l.Info("Starting server...")
		if !graph.Connect() {
			os.Exit(1)
		}
		l.Info("Connected to database!")
		if !graph.VerifySchema(false) {
			os.Exit(1)
		}
		l.Debug("Schema verified")
	},
}

func init() {
	cmd.AddCommand(c)
}

// Init causes the package to load
func Init() {}
