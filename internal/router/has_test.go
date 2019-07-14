package router

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testHas(t *testing.T, rtr Router, alias string, shouldHave bool) {
	doesHave, err := rtr.Has(alias)
	require.NoError(t, err)
	require.Equal(t, shouldHave, doesHave)
}
