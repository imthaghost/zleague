package models

// Best represents selective stats about the best x matches during the tournament time
type Best struct {
	CombinedPoints  int     `json:"combined_points"` // kills + placement
	Games           []Match `json:"games"`           // best games
	Kills           int     `json:"kills"`           // best x games total kills
	Deaths          int     `json:"deaths"`          // best x games total deaths
	KD              float64 `json:"kd"`              // best x games overall KD
	DamageDone      int     `json:"damage_done"`     // best x games total damage
	DamageTaken     int     `json:"damage_taken"`
	Headshots       int     `json:"headshots"`
	Wins            int     `json:"wins"`
	PlacementPoints int     `json:"placement"` // placement points only
}

// Total represents total stats about all matches during the tournament time
type Total struct {
	GamesPlayed    int     `json:"games_played"`
	Kills          int     `json:"kills"`
	Deaths         int     `json:"deaths"`
	Assists        int     `json:"assists"`
	Headshots      int     `json:"headshots"`
	KD             float64 `json:"kd"`
	DamageDone     int     `json:"damage_done"`
	DamageTaken    int     `json:"damage_taken"`
	Wins           int     `json:"wins"`
	CombinedPoints int     `json:"combined_points"`
}
