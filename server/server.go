package server

import (
	"zleague/api/db"
	"zleague/api/tournament"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server is a wrapper around our core
type Server struct {
	e       *echo.Echo
	db      *mongo.Database
	manager *tournament.TournamentManager
}

// NewServer will create a new instance of the server.
func NewServer(database *mongo.Database) *Server {
	if database == nil {
		database = db.Connect()
	}

	// create and start the tournament manager
	manager := tournament.NewTournamentManager(database)
	manager.Start()

	return &Server{
		e:       echo.New(), // new echo server to server the api
		db:      database,   // mongo database to store stuff
		manager: manager,    //lkajs;dlfkjas;dlfkjals;djfkalkasl;kjfa
	}
}

// GetDB will return the database connection
func (s *Server) GetDB() *mongo.Database {
	return s.db
}

// GetManager will return the tournament manager that is on the server
func (s *Server) GetManager() *tournament.TournamentManager {
	return s.manager
}

// Start will start the server instance
func (s *Server) Start(port string) {
	// default port 8080
	if port == "" {
		port = ":8080"
	}

	// register routes
	s.Routes()

	s.e.Logger.Fatal(s.e.Start(port))
}

// Stop will stop the server
func (s *Server) Stop() {
	// stop the server
	s.e.Close()
}
