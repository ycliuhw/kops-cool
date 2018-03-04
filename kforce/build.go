package kforce

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Build - Command to build configuration/templates for a new or existing cluster
type Build struct {
	requiredPaths []string
}

func (c *Build) exec() error {
	fmt.Print("this is Build.exec()!")
	return nil
}

// NewCmdBuild -
func NewCmdBuild(s *State) *cobra.Command {
	var buildCmd = &cobra.Command{
		Use:   "build",
		Short: "kforce build",
		Long:  `Deploy...`,
		Run: func(cmd *cobra.Command, args []string) {
			c := Build{}
			if err := c.exec(); err != nil {
				exitWithError(err)
			}
		},
	}
	// add additional flags for Cmd New
	// ...
	return buildCmd
}
