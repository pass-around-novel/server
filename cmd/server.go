package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var server = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "Runs the server",
	Long:    "Starts the Pass Around Novel API server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting server...")
	},
}

func init() {
	root.AddCommand(server)
}
