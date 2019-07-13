package pathtable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testSet(t *testing.T, client Client, project ProjectInfo) {
	err := client.Set(testLocation, project)
	require.NoError(t, err)
}
