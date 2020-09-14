package server

import (
	"crypto/subtle"
	"zleague/api/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) Routes() {
	// Logging
	s.e.Use(middleware.Logger())
	// CORS
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PATCH},
	}))
	// Basic Auth
	s.e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("backend")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("hackerman")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	// create a new handler and dependency inject the db and the tournament manager
	r := handlers.NewHandler(s.db, s.manager)

	// Heya!
	s.e.GET("/", r.Hello)

	// TODO: protect with basic "password" as auth header or something
	// Create Tournament
	s.e.POST("/tournament", r.NewTournament)

	// Get single tournament
	s.e.GET("/tournament/:id", r.GetTournament)

	// Update a single tournament

	// Get Team Data
	s.e.GET("/team/:teamname", r.GetTeam)

	// Update Team Data

	// Get All Teams
	s.e.GET("/teams/:id", r.GetTeams)

	// Verify Player
	s.e.GET("/check/:id", r.Verify)
}
