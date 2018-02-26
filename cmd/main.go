package main

import (
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

	fmt.Printf("\nHello, %s! Kforce is here ✌️  ✌️  ✌️\n%s\n\n", Author, strings.Repeat("-", 100))

	argv := new(kforce.CmdArgs)

	fmt.Printf("Doing -> %s\n", argv.Do())
}
