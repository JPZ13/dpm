package uninstall

import (
	"fmt"
	"os"

	"github.com/JPZ13/dpm/cmd/install"
	"github.com/JPZ13/dpm/internal/parser"
	"github.com/JPZ13/dpm/internal/project"
	"github.com/JPZ13/dpm/internal/utils"
)

// LegacyUninstallCommand houses the old logic
// for uninstalling commands
func LegacyUninstallCommand(args []string) {
	// TODO: add option to remove images
	// will need to implement something that shows what
	// images are used by each project
	if len(args) == 0 {
		uninstallAll()
	} else {
		err := uninstallListedPackages(args)
		utils.HandleFatalError(err)

		err = install.YAMLPackages()
		utils.HandleFatalError(err)
	}
}

func uninstallAll() {
	err := os.RemoveAll(project.ProjectCmdPath)
	utils.HandleFatalError(err)

	fmt.Println("All commands uninstalled")
}

func uninstallListedPackages(packages []string) error {
	commands := parser.GetCommands(project.ProjectFilePath)

	for _, pkg := range packages {
		if _, ok := commands[pkg]; ok {
			delete(commands, pkg)
			continue
		}
		return fmt.Errorf("Command %s not in project", pkg)
	}

	return parser.UpdateCommands(project.ProjectFilePath, commands)
}