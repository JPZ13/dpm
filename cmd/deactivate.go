package cmd

import (
	"github.com/JPZ13/dpm/cmd/deactivate"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(deactivateCmd)
}

var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivates the project in the current shell",
	Run: func(cmd *cobra.Command, args []string) {
		deactivate.LegacyDeactivateCommand()
	},
}
