package activate

import (
	"fmt"

	"github.com/JPZ13/dpm/cmd/install"
	"github.com/JPZ13/dpm/internal/alias"
	"github.com/JPZ13/dpm/internal/project"
	"github.com/JPZ13/dpm/internal/shell"
	"github.com/JPZ13/dpm/internal/utils"
)

// LegacyActivateCommand is the old activate command
// it is to be removed when the new activate logic
// is put in place
func LegacyActivateCommand() {
	if !project.IsProjectInstalled() {
		err := install.YAMLPackages()
		utils.HandleFatalError(err)
	}

	err := project.ActivateProject()
	utils.HandleFatalError(err)

	err = alias.SetAliases()
	utils.HandleFatalError(err)

	err = shell.StartShell(shell.Activate)
	utils.HandleFatalError(err)

	fmt.Printf("Project '%s' activated\n", project.ProjectName)
}
