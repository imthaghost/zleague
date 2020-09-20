package handlers_test

import (
	"os"
	"testing"
	"zleague/api/server"
	"zleague/api/tests"
)

var (
	s *server.Server
)

// TestMain is an entrypoint into our handler tests
func TestMain(m *testing.M) {
	// setup test deps
	s = tests.SetupTestServer()
	code := m.Run()
	// do teardown here
	os.Exit(code)
}
