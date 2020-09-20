package handlers

import (
	"html"
	"net/http"
	"zleague/api/cod"

	"github.com/labstack/echo/v4"
)

// Verify is a route that takes a user ID and returns if that user id is a valid Activision ID
func (h *Handler) Verify(c echo.Context) error {
	id := html.EscapeString(c.Param("id"))

	// see if the given id is a valid activision id
	valid := cod.IsValid(id)
	exist := map[string]bool{"exist": valid}

	return c.JSON(http.StatusOK, exist)
}
