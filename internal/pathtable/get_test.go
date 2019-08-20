package pathtable

import (
	"testing"

	"github.com/JPZ13/dpm/internal/model"
	"github.com/stretchr/testify/require"
)

func testGet(t *testing.T, client Client, project *model.ProjectInfo) {
	info, err := client.Get(testLocation)
	require.NoError(t, err)
	require.Equal(t, project, info)
}
