package router

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testAdd(t *testing.T, rtr Router, alias string) {
	err := rtr.Add(alias)
	require.NoError(t, err)
}
