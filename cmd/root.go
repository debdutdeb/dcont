package cmd

import "github.com/spf13/cobra"

var rootCommand = cobra.Command{
	Use:   "dcont",
	Short: "Start a devcontainer (only gopls available now)",
}

func Execute() error {
	// add subcommands
	rootCommand.AddCommand(environmentCmd())
	return rootCommand.Execute()
}
