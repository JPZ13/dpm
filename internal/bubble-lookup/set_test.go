package bubblelookup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetAliasInfo(t *testing.T) {
	client := makeTestClient()

	goAlias := makeGoAliasInfo()

	err := client.Set(testLocation, goAlias)
	require.NoError(t, err)

	removeTestLocationFile(t)
}
