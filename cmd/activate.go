package cmd

import (
	"github.com/JPZ13/dpm/cmd/activate"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(activateCmd)
}

var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activates the project in the current shell",
	Run: func(cmd *cobra.Command, args []string) {
		activate.LegacyActivateCommand()
	},
}
