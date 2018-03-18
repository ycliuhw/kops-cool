package kforce

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// Diff - Command to Diff configuration/templates
type Diff struct {
	*State

	requiredPaths []string
}

func (c *Diff) exec(s *State) error {
	fmt.Print("this is Diff.exec()!")
	return nil
}

func (c *Diff) getRequiredPaths(s *State) []string {
	return append(s.requiredPaths, c.requiredPaths...)
}

// NewCmdDiff -
func NewCmdDiff(s *State) *cobra.Command {
	var diffCmd = &cobra.Command{
		Use:   "diff",
		Short: "kforce diff",
		Long:  `Deploy...`,
		Run: func(cmd *cobra.Command, args []string) {
			c := &Diff{}
			fmt.Printf("diffCmd: args -> %+v\n", strings.Join(args, " "))
			fmt.Printf("diffCmd: state -> %s\n", s)
			if err := BuildCMD(c)(s); err != nil {
				exitWithError(err)
			}
		},
	}
	// add additional flags for Cmd New
	// ...
	return diffCmd
}
