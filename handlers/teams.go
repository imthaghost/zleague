package handlers

import (
	cmap "github.com/orcaman/concurrent-map"
	"html"
	"net/http"
	"strings"
	"zleague/api/models"
	"zleague/api/proxy"
	"zleague/api/tournament"

	"github.com/labstack/echo/v4"
)

// GetTeams returns all teams from a tournament
func (h *Handler) GetTeams(c echo.Context) error {
	tournamentID := html.EscapeString(c.Param("id"))

	teams, err := h.manager.GetTeams(tournamentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "tournament not found or error occurred when finding tournament")
	}

	return c.JSON(http.StatusOK, teams)
}

// GetTeam returns a single team from database
func (h *Handler) GetTeam(c echo.Context) error {
	tournamentID := html.EscapeString(c.Param("id"))
	name := html.EscapeString(c.Param("teamname"))

	team, err := h.manager.GetTeam(tournamentID, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "team in that tournament with that name not found")
	}

	return c.JSON(http.StatusOK, team)
}

// createTeamPayload will create a new team in the database
type createTeamPayload struct {
	TournamentID string `json:"tournament_id"` // id of the tournament oadd the team to
	Name string `json:"name"` // the name of the team
	Division string `json:"division"`
	Players []string `json:"players"`
}

// CreateTeam will will add a team to an already existing tournament
func (h *Handler) CreateTeam(c echo.Context) error {
	payload := createTeamPayload{}
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error binding team payload to json struct")
	}

	// clean up player names to remove # and replace with %23 cause tracker is dumb
	var cleanedNames []string
	for _, team := range payload.Players {
		cleanedNames = append(cleanedNames, strings.Replace(team, "#", "%23", 1))
	}

	tBasic := models.TeamBasic{
		Teamname:  payload.Name,
		Division:  payload.Division,
		Teammates: cleanedNames,
	}

	client := proxy.NewNetClient()
	blocked := cmap.New()

	team, err := tournament.CreateTeam(tBasic, client, &blocked)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// insert the team into the db
	t := models.Tournament{ID: payload.TournamentID}
	t.AddTeam(h.db, team)

	return c.JSON(200, team)
}

// updateTeamPayload allows us to update certain parts about a team
type updateTeamPayload struct {

}

// UpdateTeam will update a single team (mainly to correct bad stats)
func (h *Handler) UpdateTeam(c echo.Context) error {
	return nil
}

// GetTeamsByDivision returns all the teams for the given division
func (h *Handler) GetTeamsByDivision(c echo.Context) error {
	tournamentID := html.EscapeString(c.Param("id"))
	division := html.EscapeString(c.Param("div"))

	teams, err := h.manager.GetTeamsByDivision(tournamentID, division)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "teams for that division not found")
	}

	return c.JSON(http.StatusOK, teams)
}
