package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Team represents a single team in the tournament
type Team struct {
	Teamname        string   `json:"team_name"`
	Division        string   `json:"division"`
	Kills           int      `json:"kills"`
	Deaths          int      `json:"deaths"`
	Assists         int      `json:"assists"`
	KD              float64  `json:"kd"`
	GamesPlayed     int      `json:"games_played"`
	PlacementPoints int      `json:"placement_points"`
	TotalPoints     int      `json:"total_points"`
	Headshots       int      `json:"headshots"`
	DamageDone      int      `json:"damage_done"`
	LastMatch       string   `json:"last_match"`
	Wins            int      `json:"wins"`
	Players         []Player `json:"players"`
	Total           struct {
		TotalKills     int     `json:"kills"`
		TotalDeaths    int     `json:"deaths"`
		TotalAssists   int     `json:"assists"`
		TotalHeadshots int     `json:"headshots"`
		TotalKD        float64 `json:"kd"`
		TotalDamage    int     `json:"damage"`
		TotalWins      int     `json:"wins"`
		TotalScore     int     `json:"score"`
	} `json:"total"`
}

// BasicTeam is a cleaned up version to a team
type BasicTeam struct {
	Teamname        string `json:"team_name"`
	Wins            int    `json:"wins"`
	Kills           int    `json:"kills"`
	GamesPlayed     int    `json:"games_played"`
	TotalPoints     int    `json:"total_points"`
	PlacementPoints int    `json:"placement_points"`
}

// ByPoints allows us to sort all the teams
type ByPoints []Team

func (a ByPoints) Len() int           { return len(a) }
func (a ByPoints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPoints) Less(i, j int) bool { return a[i].TotalPoints > a[j].TotalPoints }

// FindTeam finds a team with the matching team name
func (t *Team) FindTeam(db *mongo.Database, teamName string) (Team, error) {
	coll := db.Collection("teams")

	// find a team based off of their name
	err := coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "teamname", Value: teamName}}).Decode(&t)
	if err != nil {
		return Team{}, err
	}

	return *t, nil
}
