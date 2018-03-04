package kforce

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Install - Command to install cluster addons
type Install struct {
	requiredPaths []string
}

func (c *Install) exec() error {
	fmt.Print("this is Install.exec()!")
	return nil
}

// NewCmdInstall -
func NewCmdInstall(s *State) *cobra.Command {
	var installCmd = &cobra.Command{
		Use:   "install",
		Short: "kforce install",
		Long:  `Deploy...`,
		Run: func(cmd *cobra.Command, args []string) {
			c := Install{}
			if err := c.exec(); err != nil {
				exitWithError(err)
			}
		},
	}
	// add additional flags for Cmd New
	// ...
	return installCmd
}
