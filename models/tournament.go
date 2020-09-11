package models

// Tournament holds information regarding the tournament
type Tournament struct {
	TournamentName string
	Teams          []Team
	// number of best games to calculate all the scores,
	// use 0 for all games
	BestGames int
}
