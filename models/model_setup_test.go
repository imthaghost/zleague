package models_test

import (
	"os"
	"testing"
	"zleague/api/server"
	"zleague/api/tests"
)

var (
	s *server.Server
)

// TestMain is the entrypoint into our model tests
func TestMain(m *testing.M) {
	// do setup here
	s = tests.SetupTestServer()
	code := m.Run()
	// do teardown here
	os.Exit(code)
}
