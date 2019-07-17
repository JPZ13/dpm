package deactivate

import (
	"fmt"

	"github.com/JPZ13/dpm/internal/alias"
	"github.com/JPZ13/dpm/internal/project"
	"github.com/JPZ13/dpm/internal/shell"
	"github.com/JPZ13/dpm/internal/utils"
)

// LegacyDeactivateCommand holds the old
// deactivate logic that is being deprecated
func LegacyDeactivateCommand() {
	err := project.DeactivateProject()
	utils.HandleFatalError(err)

	err = alias.UnsetAliases()
	utils.HandleFatalError(err)

	err = shell.StartShell(shell.Deactivate)
	utils.HandleFatalError(err)

	fmt.Printf("Project '%s' deactivated\n", project.ProjectName)
}
