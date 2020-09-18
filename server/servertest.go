package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetTestContext will return a test context
func (s *Server) GetTestContext(r *http.Request, w http.ResponseWriter) echo.Context {
	return s.e.NewContext(r, w)
}