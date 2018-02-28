package kforce

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Build - Command to build configuration/templates for a new or existing cluster
type Build struct {
	Command
	requiredPaths []string
}

func (c *Build) exec() error {
	fmt.Print("this is Build.exec()!")
	return nil
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
		if err := c.exec(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	},
}
