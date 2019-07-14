package router

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	baseTestDirectory = "./test-output/router"
	outputDirectory   = "./test-output"
)

func makeTestRouter() Router {
	return NewRouter(&Config{
		BaseDirectory: baseTestDirectory,
	})
}

func cleanTestOutput(t *testing.T) {
	err := os.RemoveAll(outputDirectory)
	require.NoError(t, err)
}
