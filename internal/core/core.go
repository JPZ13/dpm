package core

import (
	"github.com/JPZ13/dpm/internal/pathtable"
	"github.com/JPZ13/dpm/internal/router"
)

// Service holds core methods
type Service interface {
	Runner
	Project
}

// Config holds service configuration
type Config struct {
	PathTable pathtable.Client
	Router    router.Router
}

type service struct {
	runner
	project
}

type baseService struct {
	pathTable pathtable.Client
	router    router.Router
}

// New inits a core service
func New(config *Config) Service {
	base := baseService{
		pathTable: config.PathTable,
		router:    config.Router,
	}

	return &service{
		runner{base},
		project{base},
	}
}
