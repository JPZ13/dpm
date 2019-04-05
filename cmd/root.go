package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is the base Cobra command for the CLI
var RootCmd = &cobra.Command{
	Use:   "dpm",
	Short: "Install development tools locally to your project using docker containers",
}
