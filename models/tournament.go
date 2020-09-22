package models

import (
	"context"
	"time"
	"zleague/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Tournament struct holds the information needed to start a tournament.
type Tournament struct {
	ID    string        `json:"id"`    // ID single string to identify a single tournament
	Teams []models.Team `json:"teams"` // A list of teams in the tournament
	Rules struct {
		StartTime    time.Time `json:"start_time"` // Start time of tournament
		EndTime      time.Time `json:"end_time"`   // End time of tournament
		BestGamesNum int       `json:"best_games_num"`
	} `json:"rules"`
}

// TeamBasic holds a simple struct of what a team consists of.
type TeamBasic struct {
	Teamname  string    `json:"team_name"`
	Teammates []string  `json:"teammates"`
	Start     time.Time `json:"start_time"`
	End       time.Time `json:"end_time"`
	Division  string    `json:"division"`
}

// Insert will add a new tournament to the database
func (t *Tournament) Insert(db *mongo.Database) error {
	coll := db.Collection("tournaments")

	_, err := coll.InsertOne(context.TODO(), t)
	if err != nil {
		return err
	}

	return nil
}

// UpdateInDB updates the teams in the database once we have updated the players
func (t *Tournament) UpdateInDB(db *mongo.Database) {
	coll := db.Collection("tournaments")

	filter := bson.M{
		"id": t.ID,
	}

	update := bson.M{
		"$set": bson.M{
			"teams": t.Teams,
		},
	}

	_ = coll.FindOneAndUpdate(context.TODO(), filter, update)
}

// GetTeams returns all teams from a single tournament
func (t *Tournament) GetTeams(db *mongo.Database, id string) ([]models.Team, error) {
	// get tournaments collection and find single tournament
	err := db.Collection("tournaments").FindOne(context.TODO(), bson.M{"id": id}).Decode(&t)
	if err != nil {
		return []models.Team{}, err
	}

	return t.Teams, nil
}

// GetTournament returns a single tournament struct
func (t *Tournament) GetTournament(db *mongo.Database, id string) (Tournament, error) {
	// get tournaments collection and find single tournament
	err := db.Collection("tournaments").FindOne(context.TODO(), bson.M{"id": id}).Decode(&t)
	if err != nil {
		return Tournament{}, err
	}

	return *t, nil
}
