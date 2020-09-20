package handlers

import (
	"html"
	"net/http"
	"time"
	"zleague/api/tournament"

	"github.com/labstack/echo/v4"
)

// TournamentPayload represents the incoming payload to create a new tournament
type TournamentPayload struct {
	ID        string `json:"id" form:"id"`
	Start     string `json:"start" form:"start"`
	End       string `json:"end" form:"end"`
	BestGames int    `json:"best_games" form:"best_games"`
}

// CreateTournament will start a new tournament.
// TODO: Allow the ability to start and end tournaments at any time, as well as be able to set best x games :)
func (h *Handler) CreateTournament(c echo.Context) (err error) {
	tournamentPayload := TournamentPayload{}

	// csv file
	csvFormFile, err := c.FormFile("csv")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid CSV Uploaded")
	}

	// attempt bind tournament payload to our form struct
	if err := c.Bind(&tournamentPayload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Tournament Payload")
	}

	// parse the start time from the json payload
	start, err := time.Parse(time.RFC3339, tournamentPayload.Start)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Start Time")
	}
	// parse the end time from the json payload
	end, err := time.Parse(time.RFC3339, tournamentPayload.End)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid End Time")
	}

	csvData, err := csvFormFile.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error Opening CSV")
	}

	// create a new tournament
	tournament := h.manager.NewTournament(start, end, tournamentPayload.ID, csvData)

	return c.JSON(http.StatusOK, tournament)
}

// GetTournament will return a tournament that is in the database
func (h *Handler) GetTournament(c echo.Context) error {
	id := html.EscapeString(c.Param("id"))

	// get the tournament (uses the cache)
	t, err := h.manager.GetTournament(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Tournament with that ID was not found.")
	}

	return c.JSON(http.StatusOK, t)
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
