package kforce

import (
	"flag"
	"fmt"
	"os"
	"reflect"
)

// CmdArgs - parameters for cmds
type CmdArgs struct {
	Env         string
	AccountName string
	VpcID       string
	Region      string
	Debug       bool
}

// Init CmdArgs
func (args CmdArgs) validate() (string, CmdArgs) {
	const REQUIRED = "REQUIRED"

	flagSet := flag.NewFlagSet("", flag.ExitOnError)

	flagSet.StringVar(&args.Env, "Env", REQUIRED, "one of [u|s|p|m]")
	flagSet.StringVar(&args.AccountName, "AccountName", REQUIRED, "aws account name")
	flagSet.StringVar(&args.VpcID, "VpcID", REQUIRED, "Vpc ID: vpc-xxxxx")
	flagSet.StringVar(&args.Region, "Region", REQUIRED, "aws region")
	flagSet.BoolVar(&args.Debug, "Debug", false, "enable debug or not")

	flagSet.Parse(os.Args[2:])

	action := os.Args[1]

	fields := reflect.TypeOf(args)
	values := reflect.ValueOf(args)
	num := fields.NumField()
	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		if value.Interface() == REQUIRED {
			panic(field.Name + " is required!!!")
		}
		// fmt.Print("Type:", field.Type, ",", field.Name, "=", value, value.Interface(), "\n")
	}
	return action, args
}

// Do - exec CMD
func (args CmdArgs) Do() string {
	// reflect.ValueOf(&args).MethodByName(action).Call([]reflect.Value{})

	action, args := args.validate()

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
	return "new Done!"
}

// build cluster template
func (args CmdArgs) build() string {
	fmt.Printf("build... with args -> %+v\n", args)
	return "build Done!"
}

// diff cluster template
func (args CmdArgs) diff() string {
	fmt.Printf("diff... with args -> %+v\n", args)
	return "diff Done!"
}

// apply cluster template
func (args CmdArgs) apply() string {
	fmt.Printf("apply... with args -> %+v\n", args)
	return "apply Done!"
}

// install cluster addons
func (args CmdArgs) install() string {
	fmt.Printf("install... with args -> %+v\n", args)
	return "install Done!"
}
