package handlers

import (
	"html"
	"log"
	"net/http"
	"time"
	"zleague/api/tournament"

	"github.com/labstack/echo/v4"
)

func (h *Handler) NewTournament(c echo.Context) (err error) {
	// try parsing start time
	start, err := time.Parse(time.RFC3339, "2020-09-11T01:50:00+00:00")
	if err != nil {
		log.Println("Line 15 NewTournament:", err)
	}
	// try parsing end time
	end, err := time.Parse(time.RFC3339, "2020-09-11T4:50:00+00:00")
	if err != nil {
		log.Println("Line 20 NewTournament:", err)
	}
	// create a new tournament
	h.manager.NewTournament(h.db, start, end, "123458")
	// resp - 200 - check ur console hoe
	return c.JSON(http.StatusOK, "check ur console hoe")
}

func (h *Handler) GetTournament(c echo.Context) error {
	// sanitize
	id := html.EscapeString(c.Param("id"))
	// instantiate tournament struct
	m := tournament.Tournament{}
	// find tournament
	tournament := m.GetTournament(h.db, id)
	// resp - 200 - OK - {tournament}
	return c.JSON(http.StatusOK, tournament)
}
