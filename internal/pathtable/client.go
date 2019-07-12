package pathtable

// Client is the interface for pathtable methods
type Client interface {
	Get(path string) ([]AliasInfo, error)
	Set(path string, info AliasInfo) error
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
