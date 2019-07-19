package cmd

import (
	"github.com/JPZ13/dpm/cmd/list"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available commands in the current project",
	Run: func(cmd *cobra.Command, args []string) {
		list.LegacyListCommand()
	},
}
