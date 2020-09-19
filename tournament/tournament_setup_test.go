package tournament_test

import (
	"os"
	"testing"
	"zleague/api/server"
	"zleague/api/tests"
)

var (
	s *server.Server
)

func TestMain(m *testing.M) {
	s = tests.SetupTestServer()
	code := m.Run()
	// do teardown here
	os.Exit(code)
}
