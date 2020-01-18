package initdb

import (
	"../cmd"
	"github.com/spf13/cobra"
)

var (
	c = &cobra.Command{
		Use:   "initdb",
		Short: "Initializes the database",
		Long:  "Creates the schemas and requires starting data in the database",
		Run: func(cmd *cobra.Command, args []string) {
			initGraphDB()
		},
	}
	graphOnly = &cobra.Command{
		Use:   "dgraph",
		Short: "Only DGraph",
		Long:  "Only initializes the DGraph database",
		Run: func(cmd *cobra.Command, args []string) {
			initGraphDB()
		},
	}
)

func init() {
	c.AddCommand(graphOnly)
	cmd.AddCommand(c)
}

// Init causes the package to load
func Init() {}
