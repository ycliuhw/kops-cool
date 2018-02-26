package main

import (
	"flag"
	"fmt"
	"os"
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

	flagSet := flag.NewFlagSet("", flag.ExitOnError)

	flagSet.StringVar(&argv.Env, "Env", "REQUIRED", "one of [u|s|p|m]")
	flagSet.StringVar(&argv.AccountName, "AccountName", "No AccountName", "aws account name")
	flagSet.StringVar(&argv.VpcID, "VpcID", "no VpcID", "Vpc ID: vpc-xxxxx")
	flagSet.StringVar(&argv.Region, "Region", "no Region", "aws region")
	flagSet.BoolVar(&argv.Debug, "Debug", false, "enable debug or not")

	flagSet.Parse(os.Args[2:])

	argv.Action = os.Args[1]

	fmt.Printf("%s... with args -> %s\n", argv.Action, argv.Do())
}
