package server

import (
	"zleague/api/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) Routes() {
	s.e.Use(middleware.Logger())
	// s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PATCH},
	}))

	// create a new handler and dependency inject the db and the tournament manager
	r := handlers.NewHandler(s.db, s.manager)

	// Heya!
	s.e.GET("/", r.Hello)

	// TODO: protect with basic "password" as auth header or something
	// Create Tournament
	s.e.POST("/tournament", r.NewTournament)
	// Get single tournament
	// Update a single tournament

	// Get Team Data
	// Update Team Data
	// Get All Teams
	s.e.GET("/teams/:id", r.GetTeams)

	// Verify Player
}
