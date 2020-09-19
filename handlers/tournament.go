package handlers

import (
	"html"
	"log"
	"net/http"
	"time"
	"zleague/api/tournament"

	"github.com/labstack/echo/v4"
)

// CreateTournament will start a new tournament.
// TODO: Allow the ability to start and end tournaments at any time, as well as be able to set best x games :)
func (h *Handler) CreateTournament(c echo.Context) (err error) {
	start, err := time.Parse(time.RFC3339, "2020-09-11T01:50:00+00:00")
	if err != nil {
		log.Println("Line 15 NewTournament:", err)
	}
	end, err := time.Parse(time.RFC3339, "2020-09-11T4:50:00+00:00")
	if err != nil {
		log.Println("Line 20 NewTournament:", err)
	}
	// create a new tournament
	h.manager.NewTournament(start, end, "123458")

	return c.JSON(http.StatusOK, "check ur console hoe")
}

// GetTournament will return a tournament that is in the database
func (h *Handler) GetTournament(c echo.Context) error {
	id := html.EscapeString(c.Param("id"))
	m := tournament.Tournament{}

	tournament := m.GetTournament(h.db, id)

	return c.JSON(http.StatusOK, tournament)
}

// UpdateTournament will update the tournament with the given body
func (h *Handler) UpdateTournament(c echo.Context) error {
	id := html.EscapeString(c.Param("id"))

	t := tournament.Tournament{}
	// bind the body to the tournament struct
	if err := c.Bind(&t); err != nil {
		return err
	}
	// set the id because we do not want to be able to update the ID and we need it so that we know which model to update
	t.ID = id

	// TODO: Finish this when we get a response back from Omar

	return nil
}
