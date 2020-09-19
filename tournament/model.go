package tournament

import (
	"context"
	"time"
	"zleague/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Tournament struct holds the information needed to start a tournament.
type Tournament struct {
	ID        string        // ID single string to identify a single tournament
	StartTime time.Time     // Start time of tournament
	EndTime   time.Time     // End time of tournament
	Teams     []models.Team // A list of teams in the tournament
}

// TeamBasic holds a simple struct of what a team consists of.
type TeamBasic struct {
	Teamname  string
	Teammates []string
	Start     time.Time
	End       time.Time
	Division  string
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

// UpdateInDB updates the teams in the database once we have updates the players
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
func (t *Tournament) GetTeams(db *mongo.Database, id string) []models.Team {
	// get tournaments collection and find single tournament
	db.Collection("tournaments").FindOne(context.TODO(), bson.M{"id": id}).Decode(&t)

	return t.Teams
}

// GetTournament returns a single tournament struct
func (t *Tournament) GetTournament(db *mongo.Database, id string) Tournament {
	// get tournaments collection and find single tournament
	db.Collection("tournaments").FindOne(context.TODO(), bson.M{"id": id}).Decode(&t)

	return *t
}
