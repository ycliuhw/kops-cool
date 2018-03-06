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

func (c New) exec(s *State) error {
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

func (c New) getRequiredPaths(s *State) []string {
	return c.requiredPaths
}

// NewCmdNew -
func NewCmdNew(s *State) *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "new",
		Short: "kforce new",
		Long:  `Deploy...`,
		Run: func(cmd *cobra.Command, args []string) {
			var c SubCMD

			new := New{requiredPaths: []string{s.DirTemplate, s.DirAddon, s.clusterTemplatePath}}
			c = new
			fmt.Printf("newCmd: args -> %+v\n", strings.Join(args, " "))
			fmt.Printf("newCmd: state -> %s\n", s)
			if err := BuildCMD(c)(s); err != nil {
				exitWithError(err)
			}
		},
	}
	// add additional flags for Cmd New
	// ...
	return newCmd
}
