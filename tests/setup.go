package tests

import (
	"zleague/api/db"
	"zleague/api/server"
)

// SetupTestServer will conntect to the test database and then return the resulting server struct
func SetupTestServer() *server.Server {
	// connect to db
	d := db.ConnectTest()
	// connect to server
	return server.New(d)
}
