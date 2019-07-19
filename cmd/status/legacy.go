package status

import (
	"fmt"

	"github.com/JPZ13/dpm/internal/project"
	"github.com/JPZ13/dpm/internal/utils"
)

// LegacyStatusCommand holds the old
// status command logic
func LegacyStatusCommand() {
	isActive, err := project.IsProjectActive()
	utils.HandleFatalError(err)

	msg := fmt.Sprintf("Project is active at %s", project.ProjectFilePath)
	if !isActive {
		msg = "No project active"
	}

	fmt.Println(msg)
}
