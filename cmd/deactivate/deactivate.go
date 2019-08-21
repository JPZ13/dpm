package deactivate

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/JPZ13/dpm/internal/core"
	"github.com/JPZ13/dpm/internal/pathtable"
	"github.com/JPZ13/dpm/internal/router"
)

// Command houses the deactivate command
func Command(args []string) {
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
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("DPM error: %s", err)
	}

	dpmLocation := path.Join(pwd, "dpm.yml")

	err = core.DeactivateProject(ctx, dpmLocation)
	if err != nil {
		log.Fatalf("DPM error: %s", err)
	}
}
