package pathtable

import (
	"encoding/json"
	"errors"
	"path"

	"github.com/JPZ13/dpm/internal/model"
	"github.com/JPZ13/dpm/internal/utils"
)

// ErrNotFound signals that the project was not found
var ErrNotFound = errors.New("Project info file not found")

// Get retrieves blobs given an input path
func (c *client) Get(location string) (*model.ProjectInfo, error) {
	err := c.ensureBaseDirectory()
	if err != nil {
		return nil, err
	}

	// bubble up path checking for matches in basedirectory
	var digest *string
	for location != "/" {
		digest, err = c.getDigestFromPath(location)
		if err != nil {
			return nil, err
		}

		hasFile, _ := utils.DoesFileExist(*digest)
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

		hasFile, _ := utils.DoesFileExist(*digest)
		if !hasFile {
			return nil, ErrNotFound
		}
	}

	return getProjectInfoAtPath(*digest)
}

func getProjectInfoAtPath(location string) (*model.ProjectInfo, error) {
	bytes, err := utils.GetFileBytes(location)
	if err != nil {
		return nil, err
	}

	var projectInfo model.ProjectInfo
	err = json.Unmarshal(bytes, &projectInfo)
	if err != nil {
		return nil, err
	}

	return &projectInfo, nil
}

// GetDigest gets the hashed filename used to store project config
func (c *client) GetDigest(location string) (*string, error) {
	return c.getDigestFromPath(location)
}
