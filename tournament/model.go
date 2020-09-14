package tournament

import (
	"context"
	"time"
	"zleague/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Tournament struct holds the information needed to start a tournament.
// TeamMates is an array of Activision Usernames.
type Tournament struct {
	ID        string
	StartTime time.Time
	EndTime   time.Time
	Teams     []models.Team
}

// TeamBasic holds a simple struct of what a team consists of.
type TeamBasic struct {
	Teamname  string
	Teammates []string
	Start     time.Time
	End       time.Time
	Division  string
}

// Insert does...
func (t *Tournament) Insert(db *mongo.Database) error {
	coll := db.Collection("tournaments")

	_, err := coll.InsertOne(context.TODO(), t)
	if err != nil {
		return err
	}
	return nil
}

// UpdateInDB does...
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
	// return all teams from tournament
	return t.Teams
}

// GetTournament returns a single tournament struct
func (t *Tournament) GetTournament(db *mongo.Database, id string) Tournament {
	// get tournaments collection and find single tournament
	db.Collection("tournaments").FindOne(context.TODO(), bson.M{"id": id}).Decode(&t)
	// return tournament
	return *t
}
