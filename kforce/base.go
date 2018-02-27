package kforce

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	// Author is very important
	Author string = `Yang Kelvin Liu`
	// Version is very important as well
	Version string = `0.0.1`
)

var rootCmd = &cobra.Command{
	Use:   "kforce",
	Short: "Kforce is tool for creating and managing k8s cluster using kops templates",
	Long: `Kforce is tool for creating and managing k8s cluster using kops templates!
		Complete document is here ->
		https://github.com/ycliuhw/kops-cool`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Printf("\nHello, %s! Kforce is here ✌️  ✌️  ✌️\n%s\n\n", Author, strings.Repeat("-", 100))
	},
}

// Execute -
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Command - parent Command
type Command struct {
	Env           string
	AccountName   string
	VpcID         string
	Region        string
	Debug         bool
	requiredPaths []string
}

func (c *Command) preHook() {

}

func (c *Command) postHook() {

}

// Do -
func (c *Command) Do() {
	// prepare
	c.preHook()
	// do whatever u like
	c.exec()
	// clean up
	c.postHook()
}

func (c *Command) exec() {
	panic("this is Command.exec()! It has to be overloaded!!!")
}

func getRequiredPaths(c Command) []string {
	return c.requiredPaths
}

// // Command - interface
// type Command interface {
// 	getRequiredPaths() []string
// }
