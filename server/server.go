package server

import (
	"zleague/api/db"
	"zleague/api/tournament"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	e       *echo.Echo
	db      *mongo.Client
	manager *tournament.TournamentManager
}

// NewServer will create a new instance of the server.
func NewServer(database *mongo.Client) *Server {
	if database == nil {
		database = db.Connect()
	}

	return &Server{
		e:  echo.New(),
		db: database,
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
