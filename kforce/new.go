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

func (c *New) exec() {
	fmt.Print("this is New.exec()!")
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
		c.exec()
	},
}
