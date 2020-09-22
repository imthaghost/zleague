package models

import (
	"time"
)

// Player represents a single player in the tournament
type Player struct {
	Username   string    `json:"username"`
	Teamname   string    `json:"team_name"`
	TStartTime time.Time `json:"start_time"` // tournament end time
	TEndTime   time.Time `json:"end_time"`   // tournament start time
	LastMatch  string    `json:"-"`
	Matches    []Match   `json:"matches"`
	Total      Total     `json:"total"` // this struct holds the total games played
	Best       Best      `json:"best"`  // this struct holds the best x games as determined when a tournament is created
	Meta       Meta      `json:"meta"`  // meta represents random meta data like leven and icon
}

// Meta represents meta data about the player
type Meta struct {
	Level     int    `json:"level"`
	LevelIcon string `json:"level_icon"`
	Avatar    string `json:"avatar"`
}
