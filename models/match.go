package models

import "time"

// Match does
type Match struct {
	ID         string
	Mode       string
	StartTime  time.Time
	Kills      int
	Deaths     int
	Assists    int
	Headshots  int
	KD         float64
	TimePlayed string
	Placement  int
	DamageDone int
	Score      int
}

type ByScore []Match

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score+a[i].Kills > a[j].Kills+a[j].Score }
