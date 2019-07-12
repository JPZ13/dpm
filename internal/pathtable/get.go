package pathtable

import (
	"encoding/json"
	"errors"
	"path"

	"github.com/JPZ13/dpm/internal/utils"
)

// Get retrieves blobs given an input path
func (c *client) Get(location string) ([]AliasInfo, error) {
	// bubble up path checking for matches in basedirectory
	var digest string
	for location != "/" {
		digest, err := c.getDigestFromPath(location)
		if err != nil {
			return nil, err
		}

		hasFile, _ := utils.DoesFileExist(digest)
		if hasFile {
			break
		}

		location = path.Dir(location)
	}

	if location == "/" {
		digest, err := c.getDigestFromPath(location)
		if err != nil {
			return nil, err
		}

		hasFile, _ := utils.DoesFileExist(digest)
		if !hasFile {
			return nil, errors.New("Alias info file not found")
		}
	}

	return getAliasInfoAtPath(digest)
}

func getAliasInfoAtPath(location string) ([]AliasInfo, error) {
	bytes, err := utils.GetFileBytes(location)
	if err != nil {
		return nil, err
	}

	var aliases []AliasInfo
	err = json.Unmarshal(bytes, &aliases)
	if err != nil {
		return nil, err
	}

	return aliases, nil
}
