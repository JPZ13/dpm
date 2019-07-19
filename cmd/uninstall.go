package cmd

import (
	"github.com/JPZ13/dpm/cmd/uninstall"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(uninstallCmd)
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls all commands for the current project",
	Run: func(cmd *cobra.Command, args []string) {
		uninstall.LegacyUninstallCommand(args)
	},
}
