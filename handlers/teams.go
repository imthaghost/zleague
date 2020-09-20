package handlers

import (
	"html"
	"log"
	"net/http"
	"zleague/api/models"
	"zleague/api/tournament"

	"github.com/labstack/echo/v4"
)

// GetTeams returns all teams from a tournament
func (h *Handler) GetTeams(c echo.Context) error {
	tournamentID := html.EscapeString(c.Param("id"))
	m := tournament.Tournament{}

	// get teams in the tournament
	teams, err := m.GetTeams(h.db, tournamentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "team not found or error occured when finding team")
	}

	return c.JSON(http.StatusOK, teams)
}

// GetTeam returns a single team from database
func (h *Handler) GetTeam(c echo.Context) error {
	name := html.EscapeString(c.Param("teamname"))
	m := models.Team{}

	// bind data to the team struct and return
	team, err := m.FindTeam(h.db, name)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusOK, team)
}
