package kforce

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

const (
	s = "s"
	u = "u"
	p = "p"
	m = "m"
)

// ENVS - valid environment definitions
var ENVS = [4]string{s, u, p, m}

// Command -
type Command func(s State) error

// State - state and util funcs
type State struct {
	env         string
	accountName string
	vpcID       string
	region      string
	debug       bool

	clusterName          string
	stateStoreName       string
	stateStoreURI        string
	templateRenderedPath string
	currentVarsDir       string
	currentValueFilePath string
	currentIgDir         string
	currentSnippetsDir   string
	clusterSnippetsDir   string
	clusterTemplatePath  string

	requiredPaths []string
	preHooks      []Command
	postHooks     []Command

	DirRoot     string
	DirTemplate string
	DirAddon    string
	DirTmp      string
}

func (s *State) String() string {
	prettyStr := "\n"
	fields := reflect.TypeOf(*s)
	values := reflect.ValueOf(*s)
	num := fields.NumField()
	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		prettyStr += fmt.Sprintf("\t%s => %+v\n", field.Name, value)
	}
	return prettyStr
}

func isEnvValid(env string) error {
	for _, e := range ENVS {
		if env == e {
			return nil
		}
	}
	return fmt.Errorf("error: invalid env -> `%s` has to be in %v", env, ENVS)
}

// InitializeState - State construct
func InitializeState(state *State) error {

	if state.region == "" {
		state.region = "ap-southeast-2"
	}

	if err := isEnvValid(state.env); err != nil {
		return err
	}

	state.requiredPaths = []string{state.DirRoot, state.DirTemplate, state.DirAddon, state.DirTmp}

	state.DirRoot, _ = os.Getwd()
	state.DirTemplate = filepath.Join(state.DirRoot, "templates")
	state.DirAddon = filepath.Join(state.DirTemplate, "addons")
	state.DirTmp = filepath.Join(state.DirRoot, "tmp")

	state.clusterName = fmt.Sprintf("%s-%s.k8s.local", state.env, state.accountName)
	state.stateStoreName = fmt.Sprintf("%s-k8s-state-store", state.accountName)
	state.stateStoreURI = fmt.Sprintf("s3://%s", state.stateStoreName)

	state.templateRenderedPath = filepath.Join(state.DirRoot, "__generated__", fmt.Sprintf("%s-%s.yaml", state.env, state.accountName))
	state.currentVarsDir = filepath.Join(state.DirRoot, "vars", state.accountName)
	state.currentValueFilePath = filepath.Join(state.currentVarsDir, fmt.Sprintf("%s.yaml", state.env))
	state.currentIgDir = filepath.Join(state.currentVarsDir, fmt.Sprintf("%s-ig", state.env))
	state.currentSnippetsDir = filepath.Join(state.currentVarsDir, fmt.Sprintf("%s-snippets", state.env))
	state.clusterSnippetsDir = filepath.Join(state.DirTemplate, "snippets")
	state.clusterTemplatePath = filepath.Join(state.DirTemplate, "cluster.yaml")

	fmt.Printf("InitializeState: state -> \n%s\n", state)
	return nil
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
