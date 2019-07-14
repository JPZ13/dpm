package router

import "testing"

func TestRouter(t *testing.T) {
	rtr := makeTestRouter()

	testHas(t, rtr, "dne", false)

	testAdd(t, rtr, "go")
	testHas(t, rtr, "go", true)

	// gracefully handle second add
	testAdd(t, rtr, "go")
	testHas(t, rtr, "go", true)

	cleanTestOutput(t)
}
