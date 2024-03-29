package models

import (
	"time"
)

// Match does
type Match struct {
	ID          string    `json:"id"`
	Mode        string    `json:"mode"`
	StartTime   time.Time `json:"start_time"`
	Kills       int       `json:"kills"`
	Deaths      int       `json:"deaths"`
	Headshots   int       `json:"headshots"`
	KD          float64   `json:"-"`
	TimePlayed  string    `json:"time_played"`
	Placement   int       `json:"placement"`
	DamageDone  int       `json:"damage_done"`
	DamageTaken int       `json:"damage_taken"`
	WallBangs   int       `json:"wall_bangs"`
	Score       int       `json:"score"`
	Seen        int       `json:"-"`
	Checked     bool      `json:"-"`
}

// ByScore is an array of matches that allows us to return them sorted
type ByScore []Match

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score+a[i].Kills > a[j].Kills+a[j].Score }
