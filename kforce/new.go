package kforce

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// New - Command to initialize configuration/templates for a new cluster
type New struct {
	force          bool
	requiredPaths  []string
	DirRawTemplate string
}

func (c *New) exec(s *State) error {
	fmt.Print("this is New.exec()!")
	s.requiredPaths = c.requiredPaths

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	c.DirRawTemplate = path.Dir(filename)

	// do whatever u like for New
	// ...
	return nil
}

// NewCmdNew -
func NewCmdNew(s *State) *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "new",
		Short: "kforce new",
		Long:  `Deploy...`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// c := New{}
			fmt.Printf("newCmd: args -> %+v\n", strings.Join(args, " "))
			fmt.Printf("newCmd: state -> %s\n", s)
			// if err := c.exec(); err != nil {
			// 	exitWithError(err)
			// }
		},
	}
	// add additional flags for Cmd New
	// ...
	return newCmd
}
