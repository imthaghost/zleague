package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Tournament struct holds the information needed to start a tournament.
type Tournament struct {
	ID    string `json:"id"` // ID single string to identify a single tournament
	Active bool `json:"active"` // If the tournament is active
	Rules Rules  `json:"rules"` // tournament rules
	Teams []Team `json:"teams"` // A list of teams in the tournament
}

// Rules represents rules to do with the tournament
type Rules struct {
	TeamSize     int       `json:"team_size"`
	StartTime    time.Time `json:"start_time"`     // Start time of tournament
	EndTime      time.Time `json:"end_time"`       // End time of tournament
	BestGamesNum int       `json:"best_games_num"` // Amount of games to calculate for "best"
	GameMode     string    `json:"game_modes"` // the game mode that we want to track
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
func (t *Tournament) GetTeams(db *mongo.Database, id string) ([]Team, error) {
	// get tournaments collection and find single tournament
	err := db.Collection("tournaments").FindOne(context.TODO(), bson.M{"id": id}).Decode(&t)
	if err != nil {
		return []Team{}, err
	}

	return t.Teams, nil
}

// AddTeam will add a team to a currently running tournament.
func (t *Tournament) AddTeam(db *mongo.Database, team Team) {
	collection := db.Collection("tournaments")

	filter := bson.M{
		"id": t.ID,
	}

	update := bson.M{
		"$push": bson.M{
			"teams": team,
		},
	}

	_ = collection.FindOneAndUpdate(context.TODO(), filter, update)
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

func (t *Tournament) DeleteTournament(db *mongo.Database, id string) error {
	filter := bson.M{
		"id": id,
	}

	collection := db.Collection("tournaments")
	_, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}