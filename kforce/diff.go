package kforce

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Diff - Command to Diff configuration/templates
type Diff struct {
	requiredPaths []string
}

func (c *Diff) exec() error {
	fmt.Print("this is Diff.exec()!")
	return nil
}

// NewCmdDiff -
func NewCmdDiff(s *State) *cobra.Command {
	var diffCmd = &cobra.Command{
		Use:   "diff",
		Short: "kforce diff",
		Long:  `Deploy...`,
		Run: func(cmd *cobra.Command, args []string) {
			c := Diff{}
			if err := c.exec(); err != nil {
				exitWithError(err)
			}
		},
	}
	// add additional flags for Cmd New
	// ...
	return diffCmd
}
