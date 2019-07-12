package bubblelookup

import (
	"encoding/json"
	"os"

	"github.com/JPZ13/dpm/internal/utils"
)

// Set puts the blob at a hashed string
// of the path within the base directory
func (c *client) Set(path string, info AliasInfo) error {
	digest, err := c.getDigestFromPath(path)
	if err != nil {
		return err
	}

	// unmarshal any json at given path
	aliases, err := getAliasInfoAtPath(digest)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if aliases == nil {
		aliases = []AliasInfo{}
	}

	// append to what is currently there
	aliases = append(aliases, info)

	bytes, err := json.Marshal(aliases)
	if err != nil {
		return err
	}

	// write to hashed path address in base directory
	return utils.WriteFileBytes(digest, bytes)
}
