package pathtable

import (
	"encoding/json"
	"errors"
	"path"

	"github.com/JPZ13/dpm/internal/utils"
)

// Get retrieves blobs given an input path
func (c *client) Get(location string) (*ProjectInfo, error) {
	// bubble up path checking for matches in basedirectory
	var digest string
	var err error
	for location != "/" {
		digest, err = c.getDigestFromPath(location)
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

	return getProjectInfoAtPath(digest)
}

func getProjectInfoAtPath(location string) (*ProjectInfo, error) {
	bytes, err := utils.GetFileBytes(location)
	if err != nil {
		return nil, err
	}

	var projectInfo ProjectInfo
	err = json.Unmarshal(bytes, &projectInfo)
	if err != nil {
		return nil, err
	}

	return &projectInfo, nil
}
