package pathtable

import (
	"testing"

	"github.com/JPZ13/dpm/internal/model"
	"github.com/stretchr/testify/require"
)

func testSet(t *testing.T, client Client, project *model.ProjectInfo) {
	err := client.Set(testLocation, project)
	require.NoError(t, err)
}
