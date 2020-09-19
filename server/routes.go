package server

import (
	"crypto/subtle"
	"errors"
	"os"
	"zleague/api/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Routes will register our routes
func (s *Server) Routes() {
	// Logging
	s.e.Use(middleware.Logger())
	// CORS
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PATCH},
	}))
	// setup handler
	r := handlers.New(s.db, s.manager)

	// hello
	s.e.GET("/", r.Hello)

	/* Unprotected Tournament Routes */
	s.e.GET("/tournament/:id", r.GetTournament)

	/* Unprotected Team Routes */
	s.e.GET("/team/:teamname", r.GetTeam)
	s.e.GET("/teams/:id", r.GetTeams)

	// Verify a player exists
	s.e.GET("/check/:id", r.Verify)

	// protected routes are protected by basic auth and mostly used for admin stuff if shit hits the fan
	protected := s.e.Group("/protected")
	protected.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte(os.Getenv("SERVER_USERNAME"))) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(os.Getenv("SERVER_PASSWORD"))) == 1 {
			return true, nil
		}

		return false, errors.New("not authenticated :)")
	}))
	// create a new tournament (protected)
	protected.POST("/tournament", r.CreateTournament)

	// Update a tournament

	// Update Team Data
}
