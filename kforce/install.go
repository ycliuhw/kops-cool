package kforce

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// Install - Command to install cluster addons
type Install struct {
	requiredPaths []string
}

func (c *Install) exec(s *State) error {
	fmt.Print("this is Install.exec()!")
	return nil
}

func (c *Install) getRequiredPaths(s *State) []string {
	return append(s.requiredPaths, c.requiredPaths...)
}

// NewCmdInstall -
func NewCmdInstall(s *State) *cobra.Command {
	var installCmd = &cobra.Command{
		Use:   "install",
		Short: "kforce install",
		Long:  `Deploy...`,
		Run: func(cmd *cobra.Command, args []string) {
			c := &Install{}
			fmt.Printf("installCmd: args -> %+v\n", strings.Join(args, " "))
			fmt.Printf("installCmd: state -> %s\n", s)
			if err := BuildCMD(c)(s); err != nil {
				exitWithError(err)
			}
		},
	}
	// add additional flags for Cmd New
	// ...
	return installCmd
}
