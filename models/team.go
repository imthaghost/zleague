package models

// Team represents a single team in the tournament
type Team struct {
	Name     string   `json:"name"`
	Division string   `json:"division"`
	Players  []Player `json:"players"`
	Best     struct {
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
	Total struct {
		GamesPlayed int     `json:"games_played"`
		Kills       int     `json:"kills"`
		Deaths      int     `json:"deaths"`
		Headshots   int     `json:"headshots"`
		KD          float64 `json:"kd"`
		DamageDone  int     `json:"damage_done"`
		DamageTaken int     `json:"damage_taken"`
		Wins        int     `json:"wins"`
		Score       int     `json:"score"`
	} `json:"total"` // this struct hold the totals for ALL the games
}

// ByPoints allows us to sort all the teams
type ByPoints []Team

func (a ByPoints) Len() int           { return len(a) }
func (a ByPoints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPoints) Less(i, j int) bool { return a[i].Best.CombinedPoints > a[j].Best.CombinedPoints }
