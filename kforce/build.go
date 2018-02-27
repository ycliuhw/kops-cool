package kforce

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Build - Command to build configuration/templates for a new or existing cluster
type Build struct {
	Command
	requiredPaths []string
}

func (c *Build) exec() {
	fmt.Print("this is Build.exec()!")
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "kforce build",
	Long:  `Deploy...`,
	Run: func(cmd *cobra.Command, args []string) {
		c := Build{}
		c.exec()
	},
}
