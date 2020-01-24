package cmd

import (
	"fmt"
	"os"

	"github.com/sentadmedia/account/app"
	"github.com/sentadmedia/account/dep"
	"github.com/sentadmedia/elf/fw"
)

// NewRootCmd creates and initializes root command
func NewRootCmd(
	config app.Config,
	securityPolicy fw.SecurityPolicy,
) fw.Command {
	cmdFactory := dep.InitCommandFactory()
	startCmd := cmdFactory.NewCommand(
		fw.CommandConfig{
			Usage: "start",
			OnExecute: func(cmd *fw.Command, args []string) {
				app.Start(
					config,
					securityPolicy,
				)
			},
		},
	)

	rootCmd := cmdFactory.NewCommand(
		fw.CommandConfig{
			Usage:     "account",
			OnExecute: func(cmd *fw.Command, args []string) {},
		},
	)
	err := rootCmd.AddSubCommand(startCmd)
	if err != nil {
		panic(err)
	}
	return rootCmd
}

// Execute runs root command
func Execute(rootCmd fw.Command) {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
