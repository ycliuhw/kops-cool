package kforce

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// Build - Command to build configuration/templates for a new or existing cluster
type Build struct {
	*State

	requiredPaths []string
}

func (c *Build) exec(s *State) error {
	fmt.Println("this is Build.exec()!")
	return nil
}

func (c Build) getRequiredPaths(s *State) []string {
	return append(s.requiredPaths, c.requiredPaths...)
}

// NewCmdBuild -
func NewCmdBuild(s *State) *cobra.Command {
	var buildCmd = &cobra.Command{
		Use:   "build",
		Short: "kforce build",
		Long:  `Deploy...`,
		Run: func(cmd *cobra.Command, args []string) {
			c := &Build{}
			fmt.Printf("buildCmd: args -> %+v\n", strings.Join(args, " "))
			fmt.Printf("buildCmd: state -> %s\n", s)
			if err := BuildCMD(c)(s); err != nil {
				exitWithError(err)
			}
		},
	}
	// add additional flags for Cmd New
	// ...
	return buildCmd
}
