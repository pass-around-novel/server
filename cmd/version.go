package cmd

import (
	"github.com/spf13/cobra"
)

const (
	// MajorVersion of the server
	MajorVersion = 0
	// MinorVersion of the server
	MinorVersion = 1
	// Product name of the server
	Product = "Pass Around Novel Server"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "Prints the current version",
	Long:  "Prints the current version of the Pass Around Novel API server",
	Run: func(cmd *cobra.Command, args []string) {
		l.Infof("%s version %d.%d (%s)", Product, MajorVersion, MinorVersion, CommitID)
	},
}

func init() {
	root.AddCommand(version)
}
