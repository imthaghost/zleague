package models

import (
	"time"
)

// Player represents a single player in the tournament
type Player struct {
	TournamentStartTime time.Time `json:"start_time"`
	TournamentEndTime   time.Time `json:"end_time"`
	Total               struct {
		TotalKills     int     `json:"total_kills"`
		TotalDeaths    int     `json:"total_deaths"`
		TotalAssists   int     `json:"total_assists"`
		TotalHeadshots int     `json:"total_headshots"`
		TotalKD        float64 `json:"total_kd"`
		TotalDamage    int     `json:"total_damage"`
		TotalWins      int     `json:"total_wins"`
		TotalScore     int     `json:"total_score"`
	} `json:"total"`
	Username        string  `json:"username"`
	Wins            int     `json:"wins"`
	Kills           int     `json:"kills"`
	PlacementPoints int     `json:"placement_points"`
	Assists         int     `json:"assists"`
	Deaths          int     `json:"deaths"`
	KD              float64 `kd:"kd"`
	GamesPlayed     int     `json:"games_played"`
	DamageDone      int     `json:"damage_done"`
	Headshots       int     `json:"headshots"`
	Teamname        string  `json:"team_name"`
	Avatar          string  `json:"avatar"`
	Level           struct {
		Value     int    `json:"value"`
		LevelIcon string `json:"level_icon"`
	} `json:"level"`
	LastMatch   string  `json:"last_match"`
	Matches     []Match `json:"matches"`
	BestMatches []Match `json:"best_matches"`
}
