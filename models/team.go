package models

import "time"

// Team represents a single team in the tournament
type Team struct {
	Name     string   `json:"name"`
	Division string   `json:"division"`
	Players  []Player `json:"players"`
	Best     Best     `json:"best"`  // this struct holds the best x games as determined when a tournament is created
	Total    Total    `json:"total"` // this struct hold the totals for ALL the games
}

// TeamBasic holds a simple struct of what a team consists of.
type TeamBasic struct {
	Teamname  string    `json:"team_name"`
	Teammates []string  `json:"teammates"`
	Start     time.Time `json:"start_time"`
	End       time.Time `json:"end_time"`
	Division  string    `json:"division"`
}

// ByPoints allows us to sort all the teams
type ByPoints []Team

func (a ByPoints) Len() int           { return len(a) }
func (a ByPoints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPoints) Less(i, j int) bool { return a[i].Best.CombinedPoints > a[j].Best.CombinedPoints }
