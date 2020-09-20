package server

import (
	"crypto/subtle"
	"errors"
	"zleague/api/config"
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
	s.e.GET("/tournament/:id/team/:teamname", r.GetTeam)  // get a team by name
	s.e.GET("/tournament/:id/teams", r.GetTeams)          // get a team for the tournament
	s.e.GET("/tournament/:id/:div", r.GetTeamsByDivision) // get a team by the given division

	// Player /
	s.e.GET("/check/:id", r.Verify)
	s.e.GET("/check/teams", r.CheckTeams) // make sure that teams exist

	// protected routes are protected by basic auth and mostly used for admin stuff if shit hits the fan
	protected := s.e.Group("/protected")
	protected.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		config := config.GetAuthConfig()
		if subtle.ConstantTimeCompare([]byte(username), []byte(config.Username)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(config.Password)) == 1 {
			return true, nil
		}

		return false, errors.New("not authenticated :)")
	}))
	// create a new tournament (protected)
	protected.POST("/tournament", r.CreateTournament)

	// Update a tournament (protected)

	// Update Team Data (protected)
}
