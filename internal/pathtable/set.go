package pathtable

import (
	"encoding/json"

	"github.com/JPZ13/dpm/internal/utils"
)

// Set puts the blob at a hashed string
// of the path within the base directory
func (c *client) Set(path string, info ProjectInfo) error {
	err := c.ensureBaseDirectory()
	if err != nil {
		return err
	}

	digest, err := c.getDigestFromPath(path)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		return err
	}

	// write to hashed path address in base directory
	return utils.WriteFileBytes(digest, bytes)
}
