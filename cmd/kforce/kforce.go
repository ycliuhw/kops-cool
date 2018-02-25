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

// Init config files and dirs for a new cluster
func Init(args CmdArgs) string {
	fmt.Printf("Initing... with args -> %+v\n", args)
	return "Done!"
}

// Build cluster template
func Build(args CmdArgs) string {
	fmt.Printf("Building... with args -> %+v\n", args)
	return "Done!"
}

// Diff cluster template
func Diff(args CmdArgs) string {
	fmt.Printf("Diffing... with args -> %+v\n", args)
	return "Done!"
}

// Apply cluster template
func Apply(args CmdArgs) string {
	fmt.Printf("Applying... with args -> %+v\n", args)
	return "Done!"
}

// Install cluster addons
func Install(args CmdArgs) string {
	fmt.Printf("Installing... with args -> %+v\n", args)
	return "Done!"
}
