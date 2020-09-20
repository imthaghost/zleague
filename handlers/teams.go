package handlers

import (
	"encoding/csv"
	"html"
	"log"
	"net/http"
	"strings"
	"zleague/api/cod"

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

type invalidTeams struct {
	Invalid []string `json:"invalid"`
}

// CheckTeams returns a map of team and player who are invalid
// TODO move logic out of route
func (h *Handler) CheckTeams(c echo.Context) error {
	teamMap := map[string]invalidTeams{}
	// csv file
	csvFormFile, err := c.FormFile("csv")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid CSV Uploaded")
	}

	csvData, err := csvFormFile.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error Opening CSV")
	}

	lines, err := csv.NewReader(csvData).ReadAll()
	if err != nil {
		log.Println("cannot read from given csv")
	}

	for _, line := range lines {
		team := line[1]
		player := strings.Replace(line[2], "#", "%23", -1)
		if !cod.IsValid(player) {
			// teamname is present in map
			if _, ok := teamMap[team]; ok {
				t := teamMap[team]                    // reference map
				t.Invalid = append(t.Invalid, player) // append player to appropriate team
				teamMap[team] = t
			} else { // not present in map
				p := invalidTeams{}                   // create new struct reference
				p.Invalid = append(p.Invalid, player) // append teammate
				teamMap[team] = p                     // create key value pair
			}
		}
	}
	// empty map
	if len(teamMap) == 0 {
		return c.JSON(http.StatusOK, nil)
	}

	return c.JSON(http.StatusOK, teamMap)
}
