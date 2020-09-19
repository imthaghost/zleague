package models

import "time"

// Player represents a single player in the tournament
type Player struct {
	TournamentStartTime time.Time
	TournamentEndTime   time.Time
	Total               struct {
		TotalKills     int
		TotalDeaths    int
		TotalAssists   int
		TotalHeadshots int
		TotalKD        float64
		TotalDamage    int
		TotalWins      int
		TotalScore     int
	}
	Username        string
	Wins            int
	Kills           int
	PlacementPoints int
	Assists         int
	Deaths          int
	KD              float64
	GamesPlayed     int
	DamageDone      int
	Headshots       int
	Teamname        string
	Avatar          string
	Level           struct {
		Value     int
		LevelIcon string
	}
	LastMatch   string
	Matches     []Match
	BestMatches []Match
}
