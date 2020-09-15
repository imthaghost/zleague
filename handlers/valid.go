package handlers

import (
	"html"
	"net/http"
	"zleague/api/cod"

	"github.com/labstack/echo/v4"
)

// Verify is a route that takes a user ID and returns if that user id is a valid Activision ID
func (h *Handler) Verify(c echo.Context) error {
	// sanitize
	id := html.EscapeString(c.Param("id"))
	// does the user exist
	exist := map[string]bool{"exist": cod.IsValid(id)}
	// resp - OK - {exist: bool}
	return c.JSON(http.StatusOK, exist)
}
