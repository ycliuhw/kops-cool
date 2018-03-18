package kforce

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// Apply - Command to apply/render/deploy configuration/templates
type Apply struct {
	requiredPaths []string
}

func (c *Apply) exec(s *State) error {
	fmt.Print("this is Apply.exec()!")
	return nil
}

func (c *Apply) getRequiredPaths(s *State) []string {
	return append(c.requiredPaths, s.requiredPaths...)
}

// NewCmdApply -
func NewCmdApply(s *State) *cobra.Command {
	var applyCmd = &cobra.Command{
		Use:   "apply",
		Short: "kforce apply",
		Long: `Deploy ...
				  ...
				  ...
		`,
		Run: func(cmd *cobra.Command, args []string) {
			c := &Apply{requiredPaths: []string{s.templateRenderedPath}}
			fmt.Printf("applyCmd: args -> %+v\n", strings.Join(args, " "))
			fmt.Printf("applyCmd: state -> %s\n", s)
			if err := BuildCMD(c)(s); err != nil {
				exitWithError(err)
			}
		},
	}
	return applyCmd
}
