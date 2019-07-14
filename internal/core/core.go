package core

import (
	"github.com/JPZ13/dpm/internal/pathtable"
	"github.com/JPZ13/dpm/internal/router"
)

// Core holds core methods
type Core interface {
}

type core struct {
	pathTable       pathtable.Client
	router          router.Router
	baseDirectory   string
	routerDirectory string
}

// NewCore inits a Core instance from a Config
func NewCore(config *Config) Core {
	pathTable := pathtable.NewClient(&pathtable.Config{
		BaseDirectory: config.BaseDirectory,
	})

	rtr := router.NewRouter(&router.Config{
		BaseDirectory: config.RouterDirectory,
	})

	return &core{
		pathTable:       pathTable,
		router:          rtr,
		baseDirectory:   config.BaseDirectory,
		routerDirectory: config.RouterDirectory,
	}
}
