package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Team represents a single team in the tournament
type Team struct {
	Teamname        string
	Division        string
	Kills           int
	Deaths          int
	Assists         int
	KD              float64
	GamesPlayed     int
	PlacementPoints int
	TotalPoints     int
	Headshots       int
	DamageDone      int
	LastMatch       string
	Wins            int
	Players         []Player
	Total           struct {
		TotalKills     int
		TotalDeaths    int
		TotalAssists   int
		TotalHeadshots int
		TotalKD        float64
		TotalDamage    int
		TotalWins      int
		TotalScore     int
	}
}

// BasicTeam is a cleaned up version to a team
type BasicTeam struct {
	Teamname        string `json:"teamname"`
	Wins            int    `json:"wins"`
	Kills           int    `json:"kills"`
	GamesPlayed     int    `json:"gamesplayed"`
	TotalPoints     int    `json:"totalpoints"`
	PlacementPoints int    `json:"placementpoints"`
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
