package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnsureImageTag(t *testing.T) {
	t.Parallel()

	image := ensureImageTag("golang")
	require.Equal(t, "golang:latest", image)

	pythonAlpine := "python:alpine"
	withTag := ensureImageTag(pythonAlpine)
	require.Equal(t, pythonAlpine, withTag)
}
