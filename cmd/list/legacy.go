package list

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/JPZ13/dpm/internal/parser"
	"github.com/JPZ13/dpm/internal/project"
	"github.com/JPZ13/dpm/internal/utils"
)

// LegacyListCommand houses the old logic for the list
// commands command
func LegacyListCommand() {
	isActive, err := project.IsProjectActive()
	utils.HandleFatalError(err)

	if !isActive {
		log.Fatal("error: no active project - please run `dpm activate` first from your project root")
	}

	commands := parser.GetCommands(project.ProjectFilePath)
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	_, err = fmt.Fprintln(w, "COMMAND\tIMAGE\tENTRYPOINT")
	utils.HandleFatalError(err)

	for name, command := range commands {
		_, err = fmt.Fprintf(w, "%s\t%s\t%s\n", name, command.Image, command.Entrypoints)
		utils.HandleFatalError(err)
	}

	err = w.Flush()
	utils.HandleFatalError(err)
}