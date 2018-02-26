package kforce

import (
	"fmt"
)

// CmdArgs - parameters for cmds
type CmdArgs struct {
	Action      string
	Env         string
	AccountName string
	VpcID       string
	Region      string
	Debug       bool
}

// Do - exec CMD
func (args CmdArgs) Do() string {

	switch args.Action {
	case "new":
		return new(args)
	case "build":
		return build(args)
	case "diff":
		return diff(args)
	case "apply":
		return apply(args)
	case "install":
		return install(args)
	default:
		panic(args.Action + " is not supported!!!")
	}
}

// new config files and dirs for a new cluster
func new(args CmdArgs) string {
	fmt.Printf("new... with args -> %+v\n", args)
	return "Done!"
}

// build cluster template
func build(args CmdArgs) string {
	fmt.Printf("build... with args -> %+v\n", args)
	return "Done!"
}

// diff cluster template
func diff(args CmdArgs) string {
	fmt.Printf("diff... with args -> %+v\n", args)
	return "Done!"
}

// apply cluster template
func apply(args CmdArgs) string {
	fmt.Printf("apply... with args -> %+v\n", args)
	return "Done!"
}

// install cluster addons
func install(args CmdArgs) string {
	fmt.Printf("install... with args -> %+v\n", args)
	return "Done!"
}
