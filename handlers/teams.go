package handlers

import (
	"encoding/csv"
	"html"
	"log"
	"net/http"
	"strings"
	"zleague/api/cod"
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

type InvalidTeams struct {
	Invalid []string
}

// CheckTeams returns a map of team and player who are invalid
// TODO move logic out of route
func (h *Handler) CheckTeams(c echo.Context) error {
	teamMap := map[string]InvalidTeams{}
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
		log.Println("Cannot create new Reader ")
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
				p := InvalidTeams{}                   // create new struct reference
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
