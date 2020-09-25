package handlers

import (
	"net/http"
	"strings"
	"zleague/api/cod"

	"github.com/labstack/echo/v4"
)

// Verify is a route that takes a user ID and returns if that user id is a valid Activision ID
func (h *Handler) Verify(c echo.Context) error {
	id := c.Param("id")

	// replace the # with %23 for tracker api
	changed := strings.Replace(id, "#", "%23", 1)

	// see if the given id is a valid activision id
	valid := cod.IsValid(changed)
	exist := map[string]bool{"exist": valid}

	return c.JSON(http.StatusOK, exist)
}
