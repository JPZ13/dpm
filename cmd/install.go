package cmd

import (
	"github.com/JPZ13/dpm/cmd/install"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs all commands defined in dpm.yml in the current project",
	Run: func(cmd *cobra.Command, args []string) {
		install.LegacyInstallCommand(args)
	},
}
