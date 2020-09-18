package handlers_test

import (
	"os"
	"testing"
	"zleague/api/db"
	"zleague/api/server"
)

var (
	s *server.Server
)

// TestMain is an entrypoint into our handler tests
func TestMain(m *testing.M) {
	db := db.ConnectTest()
	s = server.NewServer(db)

	os.Exit(m.Run())
}
