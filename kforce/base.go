package kforce

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	// Author is very important
	Author string = `Yang Kelvin Liu`
	// Version is very important as well
	Version string = `0.0.1`
)

var rootCmd = &cobra.Command{
	Use:   "kforce ✌️  ✌️  ✌️",
	Short: "Kforce is a tool for creating and managing k8s cluster using kops yaml templates",
	Long: `Kforce is a tool for creating and managing k8s cluster using kops yaml templates!
		To get complete document, plz visit ->
		https://github.com/ycliuhw/kops-cool`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Do Stuff Here
		fmt.Println(fmt.Sprintf("\nHello, Kforce is made with 💝  by %s ✌️  ✌️  ✌️", Author))
		fmt.Println(strings.Repeat("-", 100))
		fmt.Println()

		err := InitializeState(&state)
		if err != nil {
			return err
		}
		return nil
	},
}

var state = State{}

func init() {
	rootCmd.PersistentFlags().StringVarP(&state.env, "env", "e", "", "one of [u|s|p|m]")
	rootCmd.MarkPersistentFlagRequired("env")
	rootCmd.PersistentFlags().StringVarP(&state.accountName, "account_name", "", "", "AWS account name")
	rootCmd.MarkPersistentFlagRequired("account_name")
	rootCmd.PersistentFlags().StringVarP(&state.vpcID, "vpc_id", "", "", "Vpc ID: vpc-xxxxx")
	rootCmd.MarkPersistentFlagRequired("vpc_id")
	rootCmd.PersistentFlags().StringVarP(&state.region, "region", "r", "ap-southeast-2", "AWS region")
	rootCmd.PersistentFlags().BoolVarP(&state.debug, "debug", "", false, "enable debug or not")

	rootCmd.AddCommand(NewCmdNew(&state))
	rootCmd.AddCommand(NewCmdApply(&state))
	rootCmd.AddCommand(NewCmdBuild(&state))
	rootCmd.AddCommand(NewCmdDiff(&state))
	rootCmd.AddCommand(NewCmdInstall(&state))
	rootCmd.AddCommand(versionCmd)

}

// Execute -
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exitWithError(err)
	}
}

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "\n%v\n", err)
	os.Exit(1)
}
