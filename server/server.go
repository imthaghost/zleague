package server

import (
	"net/http"
	"zleague/api/db"
	"zleague/api/tournament"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server is a wrapper around our core
type Server struct {
	e       *echo.Echo
	db      *mongo.Database
	manager *tournament.Manager
}

//New will create a new instance of the server.
func New(database *mongo.Database) *Server {
	if database == nil {
		database = db.Connect()
	}

	// create and start the tournament manager
	manager := tournament.NewManager(database)
	manager.Start()

	return &Server{
		e:       echo.New(), // new echo server to server the api
		db:      database,   // mongo database to store stuff
		manager: manager,    // tournament manager
	}
}

// GetDB will return the database connection
func (s *Server) GetDB() *mongo.Database {
	return s.db
}

// GetManager will return the tournament manager that is on the server
func (s *Server) GetManager() *tournament.Manager {
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

	// start server
	s.e.Logger.Fatal(s.e.Start(port))
}

// GetContext will return the context of the current echo server (mainly used for testing)
func (s *Server) GetContext(r *http.Request, w http.ResponseWriter) echo.Context {
	return s.e.NewContext(r, w)
}

// Stop will stop the server
func (s *Server) Stop() {
	// stop the server
	s.e.Close()
}
