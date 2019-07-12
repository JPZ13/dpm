package pathtable

import (
	"path"

	"github.com/JPZ13/dpm/internal/utils"
)

func (c *client) getDigestFromPath(location string) (string, error) {
	digest, err := utils.GetDigestJSONFilename(location)
	if err != nil {
		return "", err
	}

	fullPath := path.Join(c.baseDirectory, digest)
	return fullPath, nil
}
