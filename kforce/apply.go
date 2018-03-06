package kforce

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Apply - Command to apply/render/deploy configuration/templates
type Apply struct {
	requiredPaths []string
}

func (c Apply) exec(s *State) error {
	fmt.Print("this is Apply.exec()!")
	return nil
}

func (c Apply) getRequiredPaths(s *State) []string {
	paths := c.requiredPaths
	for _, path := range c.requiredPaths {
		paths = append(paths, path)
	}
	return paths
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
			var c SubCMD

			apply := Apply{requiredPaths: []string{s.templateRenderedPath}}
			c = apply
			if err := BuildCMD(c)(s); err != nil {
				exitWithError(err)
			}
		},
	}
	return applyCmd
}
