package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"html"
	"log"
	"net/http"
	"time"
	"zleague/api/models"
)

// TournamentPayload represents the incoming payload to create a new tournament
type TournamentPayload struct {
	ID        string `json:"id" form:"id"` // the "name" of the tournament
	Start     string `json:"start" form:"start"` // when you want the tournament to start recording stats
	End       string `json:"end" form:"end"` // when you want it to stop recording stats
	BestGames int    `json:"best_games" form:"best_games"` // how many games do we want to calculate the "best" for
	GameMode string `json:"game_mode" form:"game_mode"` // the game mode that we are tracking... duos.. trios.. etc
	TeamSize int `json:"team_size" form:"team_size"` // the size of a given team
}

// CreateTournament will start a new tournament.
func (h *Handler) CreateTournament(c echo.Context) error {
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

	csvData, err := csvFormFile.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error Opening CSV")
	}

	// create tournament rules
	rules, err := createRulesFromPayload(tournamentPayload)
	if err != nil {
		return err
	}

	// create a new tournament with the given rules
	tournament := h.manager.NewTournament(tournamentPayload.ID, rules, csvData)

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

func (h *Handler) DeleteTournament(c echo.Context) error {
	id := html.EscapeString(c.Param("id"))

	// delete tournament from db
	tournament := models.Tournament{}
	err := tournament.DeleteTournament(h.db, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	r := map[string]bool{"success": true}
	return c.JSON(http.StatusOK, r)
}

// TournamentExists will return if a tournament exists. We only do this because we do not want to return the entire tournament
func (h *Handler) TournamentExists(c echo.Context) error {
	id := html.EscapeString(c.Param("id"))

	for k := range h.manager.Tournaments.Items() {
		if id == k {
			// if we find the k (tournament name/id) then return exists
			r := map[string]bool{"exists": true}
			return c.JSON(200, r)
		}
	}

	// 404 if does not exist
	r := map[string]bool{"exists": false}
	return c.JSON(404, r)
}

// UpdateTournament will update the tournament with the given body
func (h *Handler) UpdateTournament(c echo.Context) error {
	id := html.EscapeString(c.Param("id"))

	t := models.Tournament{}
	// bind the body to the tournament struct
	if err := c.Bind(&t); err != nil {
		return err
	}
	// set the id because we do not want to be able to update the ID and we need it so that we know which model to update
	t.ID = id

	// TODO: Finish this when we get a response back from Omar

	return nil
}

// ForceUpdateLoop will call the update loop early if we need to
func (h *Handler) ForceUpdateLoop(c echo.Context) error {
	id := html.EscapeString(c.Param("id"))

	tournament, err := h.manager.GetTournament(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "could not find tournament in manager")
	}

	log.Println("Forcing update on tournament. ID:", tournament.ID)
	h.manager.Update(&tournament)
	tournament.UpdateInDB(h.db)
	log.Println("Update forced on tournament. ID:", tournament.ID)

	// update the cache
	t := models.Tournament{}
	updatedTournament, err := t.GetTournament(h.db, tournament.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to update the tournament in the cache")
	}

	// update the tournament in cache
	h.manager.Tournaments.Set(id, updatedTournament)

	return nil
}

// ForceUpdateCache will upate the cache with the data currently in the database
func (h *Handler) ForceUpdateCache(c echo.Context) error {
	id := html.EscapeString(c.Param("id"))

	// get tournament from db
	t := models.Tournament{}
	tournament, err := t.GetTournament(h.db, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "failed to find tournament with that id in the database")
	}

	// force update the cache with the tournament in the database
	h.manager.Tournaments.Set(id, tournament)

	return nil
}

// create rules from the tournament payload to clean up the function
func createRulesFromPayload(tournamentPayload TournamentPayload) (models.Rules, error) {
	// parse the start time from the json payload
	start, err := time.Parse(time.RFC3339, tournamentPayload.Start)
	if err != nil {
		return models.Rules{}, errors.New("could not parse start time")
	}
	// parse the end time from the json payload
	end, err := time.Parse(time.RFC3339, tournamentPayload.End)
	if err != nil {
		return models.Rules{}, errors.New("could not parse end time")
	}

	// default to trios
	var mode string
	if tournamentPayload.GameMode == "" {
		mode = "br_brtrios"
	} else {
		mode = tournamentPayload.GameMode
	}

	// default to team size of 3
	var teamSize int
	if tournamentPayload.TeamSize == 0 {
		teamSize = 3
	} else {
		teamSize = tournamentPayload.TeamSize
	}

	// default to best 4 games
	var bestGames int
	if tournamentPayload.BestGames == 0 {
		bestGames = 4
	} else {
		bestGames = tournamentPayload.BestGames
	}

	rules := models.Rules{
		StartTime: start,
		EndTime: end,
		TeamSize: teamSize,
		BestGamesNum: bestGames,
		GameMode: mode,
	}

	return rules, nil
}
