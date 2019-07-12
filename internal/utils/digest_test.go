package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	dummyDigestJSON = "sha256:b5a2c96250612366ea272ffac6d9744aaf4b45aacd96aa7cfcb931ee3b558259.json"

	pathDigestJSON = "sha256:e5bf2e31c996e0e28542259894057d51ff81c36f06cd33d5e10fbc8cdec33118.json"
)

func TestGetDigestJSON(t *testing.T) {
	dummyString := "dummy"
	jsonDummy, err := GetDigestJSONFilename(dummyString)
	require.NoError(t, err)
	require.Equal(t, jsonDummy, dummyDigestJSON)

	pathString := "/go/src/github.com/"
	pathDigest, err := GetDigestJSONFilename(pathString)
	require.NoError(t, err)
	require.Equal(t, pathDigest, pathDigestJSON)
}
