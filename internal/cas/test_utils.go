package cas

import (
	"os"
	"path"
	"testing"

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

func makeGoAliasInfo() AliasInfo {
	goAliases := map[string]bool{
		"go":     true,
		"golang": true,
	}

	return AliasInfo{
		Aliases:        goAliases,
		BinaryLocation: "/usr/local/bin",
		VolumeName:     "supple-leopard",
	}
}

func makePythonAliasInfo() AliasInfo {
	pythonAliases := map[string]bool{
		"python": true,
		"pip":    true,
	}

	return AliasInfo{
		Aliases:        pythonAliases,
		BinaryLocation: "/usr/local/bin",
		VolumeName:     "crouching-tiger",
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
