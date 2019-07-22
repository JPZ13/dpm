package cmd

import (
	"github.com/JPZ13/dpm/cmd/activate"
	"github.com/JPZ13/dpm/cmd/deactivate"
	"github.com/JPZ13/dpm/cmd/install"
	"github.com/JPZ13/dpm/cmd/list"
	"github.com/JPZ13/dpm/cmd/status"
	"github.com/JPZ13/dpm/cmd/uninstall"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(activateCmd)
	RootCmd.AddCommand(deactivateCmd)
	RootCmd.AddCommand(installCmd)
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(statusCmd)
	RootCmd.AddCommand(uninstallCmd)
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
		activate.LegacyActivateCommand()
	},
}

var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivates the project in the current shell",
	Run: func(cmd *cobra.Command, args []string) {
		deactivate.LegacyDeactivateCommand()
	},
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs all commands defined in dpm.yml in the current project",
	Run: func(cmd *cobra.Command, args []string) {
		install.LegacyInstallCommand(args)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available commands in the current project",
	Run: func(cmd *cobra.Command, args []string) {
		list.LegacyListCommand()
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows the current project status",
	Run: func(cmd *cobra.Command, args []string) {
		status.LegacyStatusCommand()
	},
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls all commands for the current project",
	Run: func(cmd *cobra.Command, args []string) {
		uninstall.LegacyUninstallCommand(args)
	},
}
