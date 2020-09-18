package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello from ZLeague
func (h *Handler) Hello(c echo.Context) (err error) {
	// hello :)
	res := map[string]string{"msg": "hello from zleague!"}
	// 200 - OK
	return c.JSON(http.StatusOK, res)
}
