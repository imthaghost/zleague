package models

import (
	"time"
)

// Player represents a single player in the tournament
type Player struct {
	Username            string    `json:"username"`
	Teamname            string    `json:"team_name"`
	TournamentStartTime time.Time `json:"start_time"`
	TournamentEndTime   time.Time `json:"end_time"`
	LastMatch           string    `json:"-"`
	Matches             []Match   `json:"matches"`
	Total               struct {
		Kills       int     `json:"kills"`
		Deaths      int     `json:"deaths"`
		Assists     int     `json:"assists"`
		Headshots   int     `json:"headshots"`
		KD          float64 `json:"kd"`
		DamageDone  int     `json:"damage_done"`
		DamageTaken int     `json:"damage_taken"`
		Wins        int     `json:"wins"`
		Score       int     `json:"score"`
	} `json:"total"` // this struct holds the total games played
	Best struct {
		CombinedPoints  int     `json:"combined_points"` // kills + placement
		Games           []Match `json:"games"`           // best games
		Kills           int     `json:"kills"`           // best x games total kills
		Deaths          int     `json:"deaths"`          // best x games total deaths
		KD              int     `json:"kd"`              // best x games overall KD
		DamageDone      int     `json:"damage_done"`     // best x games total damage
		DamageTaken     int     `json:"damage_taken"`
		Headshots       int     `json:"headshots"`
		Wins            int     `json:"wins"`
		PlacementPoints int     `json:"placement"` // placement points only
	} `json:"best"` // this struct holds the best x games as determined when a tournament is created
	Meta struct {
		Level     int    `json:"level"`
		LevelIcon string `json:"level_icon"`
		Avatar    string `json:"avatar"`
	} `json:"meta"` // meta represents random meta data like leven and icon
}
