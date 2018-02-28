package kforce

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Apply - Command to apply/render/deploy configuration/templates
type Apply struct {
	Command
	requiredPaths []string
}

func (c *Apply) exec() error {
	fmt.Print("this is Apply.exec()!")
	return nil
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
		if err := c.exec(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	},
}
