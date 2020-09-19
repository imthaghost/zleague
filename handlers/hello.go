package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello from ZLeague
func (h *Handler) Hello(c echo.Context) (err error) {
	res := map[string]string{"msg": "hello from zleague!"}

	return c.JSON(http.StatusOK, res)
}
