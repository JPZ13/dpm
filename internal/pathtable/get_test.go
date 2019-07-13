package pathtable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testGet(t *testing.T, client Client, project ProjectInfo) {
	info, err := client.Get(testLocation)
	require.NoError(t, err)
	require.Equal(t, project, *info)
}
