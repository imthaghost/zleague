package server

import (
	"zleague/api/db"
	"zleague/api/tournament"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

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

func (s *Server) Stop() {
	// stop the server
	s.e.Close()
}
