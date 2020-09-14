package server

import (
	"zleague/api/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (s *Server) Routes() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PATCH},
	}))

	// create a new handler and dependency inject the db and the tournament manager
	r := handlers.NewHandler(s.db, s.manager)

	// Heya!
	s.e.GET("/", r.Hello)
}
