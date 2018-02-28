package kforce

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New - Command to initialize configuration/templates for a new cluster
type New struct {
	Command
	force         bool
	requiredPaths []string
}

func (c *New) exec() error {
	fmt.Print("this is New.exec()!")
	return nil
}

func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "kforce new",
	Long:  `Deploy...`,
	Run: func(cmd *cobra.Command, args []string) {
		c := New{}
		if err := c.exec(); err != nil {
			exitWithError(err)
		}
	},
}
