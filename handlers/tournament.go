package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (h *Handler) NewTournament(c echo.Context) (err error) {
	// try parsing start time
	start, err := time.Parse(time.RFC3339, "2020-09-11T01:50:00+00:00")
	if err != nil {
		log.Fatal(err)
	}
	// try parsing end time
	end, err := time.Parse(time.RFC3339, "2020-09-11T4:50:00+00:00")
	if err != nil {
		log.Fatal(err)
	}

	// create a new tournament
	h.manager.NewTournament(h.db, start, end, "123458")

	return c.JSON(http.StatusOK, "check ur console hoe")
}
