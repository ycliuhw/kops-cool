package kforce

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Diff - Command to Diff configuration/templates
type Diff struct {
	Command
	requiredPaths []string
}

func (c *Diff) exec() {
	fmt.Print("this is Diff.exec()!")
}

func init() {
	rootCmd.AddCommand(diffCmd)
}

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "kforce diff",
	Long:  `Deploy...`,
	Run: func(cmd *cobra.Command, args []string) {
		c := Diff{}
		c.exec()
	},
}
