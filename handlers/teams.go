package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetTeams(c echo.Context) error {
	tournamentID := c.Param("id")

	t := h.manager.Tournaments[tournamentID]
	return c.JSON(http.StatusOK, t.GetTeams(h.manager.DB))
}
