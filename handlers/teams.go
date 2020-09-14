package handlers

import (
	"html"
	"log"
	"net/http"
	"zleague/api/models"
	"zleague/api/tournament"

	"github.com/labstack/echo"
)

// GetTeams returns all teams from a tournament
func (h *Handler) GetTeams(c echo.Context) error {
	// sanitize
	tournamentID := html.EscapeString(c.Param("id"))
	// instantiate tournament struct
	m := tournament.Tournament{}
	// get teams array
	teams := m.GetTeams(h.db, tournamentID)
	// resp - 200 - OK - [teams]
	return c.JSON(http.StatusOK, teams)
}

// GetTeam returns a single team from database
func (h *Handler) GetTeam(c echo.Context) error {
	// get teamname from params
	name := html.EscapeString(c.Param("teamname"))
	// instantiate team struct
	m := models.Team{}
	// bind data
	team, err := m.FindTeam(h.db, name)
	if err != nil {
		log.Println(err)
	}
	// resp - 200 - OK - team
	return c.JSON(http.StatusOK, team)
}
