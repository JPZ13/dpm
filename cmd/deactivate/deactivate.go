package deactivate

import (
	"context"
	"log"
	"os"

	"github.com/JPZ13/dpm/cmd/tools"
)

// Command houses the deactivate command
func Command(args []string) {
	core, err := tools.MakeCoreInHomeDirectory()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	ctx := context.Background()
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("DPM error: %s", err)
	}

	err = core.DeactivateProject(ctx, pwd)
	if err != nil {
		log.Fatalf("DPM error: %s", err)
	}
}
