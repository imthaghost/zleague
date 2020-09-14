package handlers

import (
	"html"
	"net/http"
	"zleague/api/cod"

	"github.com/labstack/echo"
)

func (h *Handler) Verify(c echo.Context) error {
	// sanitize
	id := html.EscapeString(c.Param("id"))
	// does the user exist
	exist := map[string]bool{"exist": cod.IsValid(id)}
	// resp - OK - {exist: bool}
	return c.JSON(http.StatusOK, exist)
}
