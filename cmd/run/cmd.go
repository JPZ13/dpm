package run

import (
	"context"
	"log"

	"github.com/JPZ13/dpm/cmd/tools"
)

// Command is the top level command
// for dpm run
func Command(args []string) {
	core, err := tools.MakeCoreInHomeDirectory()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	ctx := context.Background()
	err = core.Run(ctx, args)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
}
