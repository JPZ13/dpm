package run

import (
	"context"

	"github.com/JPZ13/dpm/internal/core"
	"github.com/JPZ13/dpm/internal/pathtable"
	"github.com/JPZ13/dpm/internal/router"
)

// Command is the top level command
// for dpm run
func Command(args []string) error {
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
	return core.Run(ctx, args)
}
