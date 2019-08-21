package tools

import (
	"os"
	"path"

	"github.com/JPZ13/dpm/internal/core"
	"github.com/JPZ13/dpm/internal/pathtable"
	"github.com/JPZ13/dpm/internal/router"
)

// MakeCoreInHomeDirectory inits a core service
// and puts the router/pathtable in the .dpm folder
// in a user's home directory
func MakeCoreInHomeDirectory() (core.Service, error) {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	baseDirectory := path.Join(homeDirectory, ".dpm")
	routerDirectory := path.Join(baseDirectory, "router")

	pt := pathtable.NewClient(&pathtable.Config{
		BaseDirectory: baseDirectory,
	})

	rtr := router.NewRouter(&router.Config{
		BaseDirectory: routerDirectory,
	})

	core := core.New(&core.Config{
		PathTable: pt,
		Router:    rtr,
	})

	return core, nil
}
