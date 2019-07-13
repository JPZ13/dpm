package router

import (
	"os"

	"github.com/JPZ13/dpm/internal/utils"
)

// Router holds public methods for a router
type Router interface {
	// Has(alias string) (bool, error)
}

type router struct {
	baseDirectory string
}

// NewRouter inits a Router interface
func NewRouter(config *Config) Router {
	return &router{
		baseDirectory: config.BaseDirectory,
	}
}

func (r *router) ensureDirectory() error {
	doesExist, err := utils.DoesFileExist(r.baseDirectory)
	if !doesExist {
		return os.Mkdir(r.baseDirectory, utils.WriteMode)
	}

	return err
}
