package cas

// Client is the interface for CAS methods
type Client interface {
	Get(path string) ([]AliasInfo, error)
	Set(path string, info AliasInfo) error
}

type client struct {
	baseDirectory string
}

// NewClient instantiates a CAS client
func NewClient(config *Config) Client {
	return &client{
		baseDirectory: config.BaseDirectory,
	}
}
