package pathtable

import "github.com/JPZ13/dpm/internal/model"

// Client is the interface for pathtable methods
type Client interface {
	Get(location string) (*model.ProjectInfo, error)
	Set(path string, info *model.ProjectInfo) error
}

type client struct {
	baseDirectory string
}

// NewClient instantiates a pathtable client
func NewClient(config *Config) Client {
	return &client{
		baseDirectory: config.BaseDirectory,
	}
}
