package run

import (
	"context"
	"log"

	"github.com/JPZ13/dpm/internal/core"
	"github.com/JPZ13/dpm/internal/pathtable"
	"github.com/JPZ13/dpm/internal/router"
)

// Command is the top level command
// for dpm run
func Command(args []string) {
	// TODO replace directories with config
	pt := pathtable.NewClient(&pathtable.Config{
		BaseDirectory: "~/.dpm",
	})

	rtr := router.NewRouter(&router.Config{
		BaseDirectory: "~/.dpm/router",
	})

	core := core.New(&core.Config{
		PathTable: pt,
		Router:    rtr,
	})

	ctx := context.Background()
	err := core.Run(ctx, args)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
}
