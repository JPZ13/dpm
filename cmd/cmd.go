package cmd

import (
	"github.com/JPZ13/dpm/cmd/activate"
	"github.com/JPZ13/dpm/cmd/deactivate"
	"github.com/JPZ13/dpm/cmd/run"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(activateCmd)
	RootCmd.AddCommand(deactivateCmd)
	RootCmd.AddCommand(runCmd)
}

// RootCmd is the base Cobra command for the CLI
var RootCmd = &cobra.Command{
	Use:   "dpm",
	Short: "Install development tools locally to your project using docker containers",
}

var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activates the project in the current shell",
	Run: func(cmd *cobra.Command, args []string) {
		activate.Command(args)
	},
}

var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivates the project in the current shell",
	Run: func(cmd *cobra.Command, args []string) {
		deactivate.Command(args)
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an alias that is defined in the dpm yaml file",
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
	Run: func(cmd *cobra.Command, args []string) {
		run.Command(args)
	},
}
