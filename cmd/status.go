package cmd

import (
	"github.com/JPZ13/dpm/cmd/status"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows the current project status",
	Run: func(cmd *cobra.Command, args []string) {
		status.LegacyStatusCommand()
	},
}
