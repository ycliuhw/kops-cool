package kforce

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Kforce",
	Long:  `All software has versions. This is Kforce's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Kforce -> %s", Version)
	},
}
