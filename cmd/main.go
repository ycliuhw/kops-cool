package main

import (
	"flag"
	"fmt"
	"strings"

	"./kforce"
)

const (
	// Author is very important
	Author string = `Yang Kelvin Liu`
	// Version is very important as well
	Version string = `0.0.1`
)

func main() {
	fmt.Printf("\nHello, %s! Kforce is here ✌️\n%s\n\n", Author, strings.Repeat("-", 100))

	argv := new(kforce.CmdArgs)
	flag.StringVar(&argv.Env, "Env", "no Env", "one of [u|s|p|m]")
	flag.StringVar(&argv.AccountName, "AccountName", "No AccountName", "aws account name")
	flag.StringVar(&argv.VpcID, "VpcID", "no VpcID", "Vpc ID: vpc-xxxxx")
	flag.StringVar(&argv.Region, "Region", "no Region", "aws region")
	flag.BoolVar(&argv.Debug, "Debug", false, "enable debug or not")
	flag.Parse()

	kforce.Init(*argv)
	kforce.Build(*argv)
	kforce.Apply(*argv)
	kforce.Install(*argv)
}
