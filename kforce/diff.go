package kforce

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Diff - Command to Diff configuration/templates
type Diff struct {
	Command
	requiredPaths []string
}

func (c *Diff) exec() error {
	fmt.Print("this is Diff.exec()!")
	return nil
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
		if err := c.exec(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	},
}
