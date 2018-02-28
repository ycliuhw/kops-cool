package kforce

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Install - Command to install cluster addons
type Install struct {
	Command
	requiredPaths []string
}

func (c *Install) exec() error {
	fmt.Print("this is Install.exec()!")
	return nil
}

func init() {
	rootCmd.AddCommand(installCmd)
}

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
