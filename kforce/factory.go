package kforce

import (
	"fmt"
	"os"
)

// Command -
type Command func(s State) error

// State - state and util funcs
type State struct {
	requiredPaths []string
	preHooks      []Command
	postHooks     []Command
}

func (s *State) addPreHooks(f Command) {
	s.preHooks = append(s.preHooks, f)
}

func (s *State) addPostHooks(f Command) {
	s.postHooks = append(s.postHooks, f)
}

func (s *State) addRequiredPath(path string) {
	s.requiredPaths = append(s.requiredPaths, path)
}

func (s *State) ensurePaths() error {
	for _, path := range s.requiredPaths {
		fmt.Printf("\tChecking path -> %s\n", path)
		p, err := os.Stat(path)
		if err != nil {
			fmt.Println(err)
			return err
		}
		switch mode := p.Mode(); {
		case mode.IsDir():
			// do directory stuff
			fmt.Printf("%s is directory", path)
		case mode.IsRegular():
			// do file stuff
			fmt.Printf("%s is file", path)
		default:
			return fmt.Errorf("Not exist -> %s", path)
		}

	}
	return nil
}

func (s *State) preRun() error {
	// ensure all required path are there
	err := s.ensurePaths()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// run pre steps
	for _, f := range s.preHooks {
		err := f(*s)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (s *State) postRun() error {

	// run pre steps
	for _, f := range s.postHooks {
		err := f(*s)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

// BuildCMD - build cmd
func BuildCMD(cmd Command) Command {
	return func(s State) error {
		err := s.preRun()
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = cmd(s)
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = s.postRun()
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
}
