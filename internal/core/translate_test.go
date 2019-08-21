package core

import (
	"testing"

	"github.com/JPZ13/dpm/internal/model"
	"github.com/stretchr/testify/require"
)

const testDPMFileLocation = "./test-data/dpm.yml"

func TestTranslateDPMFileToProjectInfo(t *testing.T) {
	dpmFile := &model.DPMFile{}
	err := parseYAMLFile(testDPMFileLocation, dpmFile)
	require.NoError(t, err)

	testDPMFile := makeTestDPMFile()
	require.Equal(t, *testDPMFile, *dpmFile)

	projectInfo, err := translateDPMFileToProjectInfo(dpmFile)
	require.NoError(t, err)
	require.NotNil(t, projectInfo)
}
