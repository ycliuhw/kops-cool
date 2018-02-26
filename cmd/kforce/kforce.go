package kforce

import (
	"fmt"
)

// CmdArgs - parameters for cmds
type CmdArgs struct {
	Env         string
	AccountName string
	VpcID       string
	Region      string
	Debug       bool
}

// Do - exec CMD
func (args CmdArgs) Do(action string) string {

	switch action {
	case "new":
		return args.new()
	case "build":
		return args.build()
	case "diff":
		return args.diff()
	case "apply":
		return args.apply()
	case "install":
		return args.install()
	default:
		panic(action + " is not supported!!!")
	}
}

// new config files and dirs for a new cluster
func (args CmdArgs) new() string {
	fmt.Printf("new... with args -> %+v\n", args)
	return "Done!"
}

// build cluster template
func (args CmdArgs) build() string {
	fmt.Printf("build... with args -> %+v\n", args)
	return "Done!"
}

// diff cluster template
func (args CmdArgs) diff() string {
	fmt.Printf("diff... with args -> %+v\n", args)
	return "Done!"
}

// apply cluster template
func (args CmdArgs) apply() string {
	fmt.Printf("apply... with args -> %+v\n", args)
	return "Done!"
}

// install cluster addons
func (args CmdArgs) install() string {
	fmt.Printf("install... with args -> %+v\n", args)
	return "Done!"
}
