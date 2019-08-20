package pathtable

import (
	"os"
	"path"
	"testing"

	"github.com/JPZ13/dpm/internal/model"
	"github.com/JPZ13/dpm/internal/utils"
	"github.com/stretchr/testify/require"
)

const (
	baseTestDirectory = "./test-data"

	testLocation = "/go/src/github.com/ip-man/wing-chung"
)

func makeTestClient() Client {
	return NewClient(&Config{
		BaseDirectory: baseTestDirectory,
	})
}

func makeGoAliasInfo() *model.AliasInfo {
	goAliases := map[string]string{
		"go":     "/usr/local/bin",
		"golang": "/usr/local/bin",
	}

	return &model.AliasInfo{
		Aliases:    goAliases,
		VolumeName: "supple-leopard",
		Image:      "golang:1.12",
	}
}

func makePythonAliasInfo() *model.AliasInfo {
	pythonAliases := map[string]string{
		"python": "/usr/local/bin",
		"pip":    "/usr/local/bin",
	}

	return &model.AliasInfo{
		Aliases:    pythonAliases,
		VolumeName: "crouching-tiger",
		Image:      "python:3",
	}
}

func makeProjectInfo() *model.ProjectInfo {
	return &model.ProjectInfo{
		IsActive: true,
		Commands: []model.AliasInfo{},
	}
}

func getTestDigestFromPath(t *testing.T, location string) string {
	digest, err := utils.GetDigestJSONFilename(location)
	require.NoError(t, err)
	return path.Join(baseTestDirectory, digest)
}

func removeTestLocationFile(t *testing.T) {
	filename := getTestDigestFromPath(t, testLocation)

	err := os.Remove(filename)
	require.NoError(t, err)
}
