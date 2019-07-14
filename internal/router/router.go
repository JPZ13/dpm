package router

// Router holds public methods for a router
type Router interface {
	Has(alias string) (bool, error)
	Add(alias string) error
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
