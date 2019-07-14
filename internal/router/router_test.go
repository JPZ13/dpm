package router

import "testing"

func TestRouter(t *testing.T) {
	rtr := makeTestRouter()

	testHas(t, rtr, "dne", false)

	cleanTestOutput(t)
}
