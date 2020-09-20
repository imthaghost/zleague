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
	s = tests.SetupTestServer()
	os.Exit(m.Run())
}
