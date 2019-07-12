package bubblelookup

// Client is the interface for bubble-lookup methods
type Client interface {
	Get(path string) ([]AliasInfo, error)
	Set(path string, info AliasInfo) error
}

type client struct {
	baseDirectory string
}

// NewClient instantiates a bubble-lookup client
func NewClient(config *Config) Client {
	return &client{
		baseDirectory: config.BaseDirectory,
	}
}
