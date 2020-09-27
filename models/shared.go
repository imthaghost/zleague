package models

// Best represents selective stats about the best x matches during the tournament time
type Best struct {
	Kills           int     `json:"kills"`  // best x games total kills
	Deaths          int     `json:"deaths"` // best x games total deaths
	Headshots       int     `json:"headshots"`
<<<<<<< HEAD
	KD              float64 `json:"-"`           // best x games overall KD
=======
	KD              float64 `json:"-"`          // best x games overall KD
>>>>>>> 147697a0a04a2ed8f25c3259d2d3d9f9214473fc
	DamageDone      int     `json:"damage_done"` // best x games total damage
	DamageTaken     int     `json:"damage_taken"`
	WallBangs       int     `json:"-"`
	Wins            int     `json:"wins"`
	PlacementPoints int     `json:"placement"`       // placement points only
	CombinedPoints  int     `json:"combined_points"` // kills + placement
<<<<<<< HEAD
	Games           []Match `json:"-"`               // best games
=======
	Games           []Match `json:"-"`           // best games
>>>>>>> 147697a0a04a2ed8f25c3259d2d3d9f9214473fc
}

// Total represents total stats about all matches during the tournament time
type Total struct {
	Kills           int     `json:"kills"`     // all kills in all games during the tournament
	Deaths          int     `json:"deaths"`    // all deaths in all games during the tournament
	Headshots       int     `json:"headshots"` // all headshots in all games during the tournametn
	KD              float64 `json:"-"`
	DamageDone      int     `json:"damage_done"`
	DamageTaken     int     `json:"damage_taken"`
	WallBangs       int     `json:"-"`
	Wins            int     `json:"wins"`
	PlacementPoints int     `json:"placement_points"`
	CombinedPoints  int     `json:"combined_points"`
	Games           []Match `json:"-"`
	GamesPlayed     int     `json:"games_played"`
}
