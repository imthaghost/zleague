package models_test

import (
	"os"
	"testing"
	"zleague/api/db"
	"zleague/api/server"
)

var (
	s *server.Server
)

// TestMain is the entrypoint into our model tests
func TestMain(m *testing.M) {
	db := db.ConnectTest()
	s = server.NewServer(db)

	os.Exit(m.Run())
}

func refreshTeamTable() {

}

func seedTeam() {

}
