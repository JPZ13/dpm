package cmd

import (
	"fmt"

	"github.com/JPZ13/dpm/internal/alias"
	"github.com/JPZ13/dpm/internal/project"
	"github.com/JPZ13/dpm/internal/shell"
	"github.com/JPZ13/dpm/internal/utils"
	"github.com/spf13/cobra"
)

var forceActivate bool

func init() {
	activateCmd.Flags().BoolVarP(&forceActivate, "force", "f", false,
		"Force activation even if another project is currently active")
	RootCmd.AddCommand(activateCmd)
}

var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activates the project in the current shell",
	Run: func(cmd *cobra.Command, args []string) {
		if !project.IsProjectInstalled() {
			err := installYAMLPackages()
			utils.HandleFatalError(err)
		}

		err := project.ActivateProject()
		utils.HandleFatalError(err)

		err = alias.SetAliases()
		utils.HandleFatalError(err)

		err = shell.StartShell(shell.Activate)
		utils.HandleFatalError(err)

		fmt.Printf("Project '%s' activated\n", project.ProjectName)
	},
}
