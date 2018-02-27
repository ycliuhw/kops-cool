package kforce

type command struct {
	Env           string
	AccountName   string
	VpcID         string
	Region        string
	Debug         bool
	requiredPaths []string
}

func (c *command) Do() {
	panic("this is command.Do()! It has to be overloaded!!!")
}

func getRequiredPaths(c command) []string {
	return c.requiredPaths
}

// Command - interface
type Command interface {
	getRequiredPaths() []string
}

// New - command to initialize configuration/templates for a new cluster
type New struct {
	command
	force         bool
	requiredPaths []string
}

// Do -
func (c *New) Do() {
	panic("this is New.Do()!")
}

// Build - command to build configuration/templates for a new or existing cluster
type Build struct {
	command
	requiredPaths []string
}

// Do -
func (c *Build) Do() {
	panic("this is Build.Do()!")
}

// Diff - command to Diff configuration/templates
type Diff struct {
	command
	requiredPaths []string
}

// Do -
func (c *Diff) Do() {
	panic("this is Diff.Do()!")
}

// Apply - command to apply/render/deploy configuration/templates
type Apply struct {
	command
	requiredPaths []string
}

// Do -
func (c *Apply) Do() {
	panic("this is Apply.Do()!")
}

// Install - command to install cluster addons
type Install struct {
	command
	requiredPaths []string
}

// Do -
func (c *Install) Do() {
	panic("this is Install.Do()!")
}
