package kforce

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Apply - Command to apply/render/deploy configuration/templates
type Apply struct {
	Command
	requiredPaths []string
}

func (c *Apply) exec() {
	fmt.Print("this is Apply.exec()!")
}

func init() {
	rootCmd.AddCommand(applyCmd)
}

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "kforce apply",
	Long:  `Deploy...`,
	Run: func(cmd *cobra.Command, args []string) {
		c := Apply{}
		c.exec()
	},
}
