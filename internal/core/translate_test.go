package core

import (
	"path"
	"testing"

	"github.com/JPZ13/dpm/internal/model"
	"github.com/JPZ13/dpm/internal/utils"
	"github.com/stretchr/testify/require"
)

const testDPMFileLocation = "./test-data/dpm.yml"

func TestTranslateDPMFileToProjectInfo(t *testing.T) {
	t.Parallel()

	dpmFile := &model.DPMFile{}
	err := parseYAMLFile(testDPMFileLocation, dpmFile)
	require.NoError(t, err)

	testDPMFile := makeTestDPMFile()
	require.Equal(t, *testDPMFile, *dpmFile)

	dpmDir := path.Dir(testDPMFileLocation)
	digest, err := utils.GetDigestJSONFilename(dpmDir)
	require.NoError(t, err)

	projectInfo, err := translateDPMFileToProjectInfo(dpmFile, digest)
	require.NoError(t, err)
	require.NotNil(t, projectInfo)
}
