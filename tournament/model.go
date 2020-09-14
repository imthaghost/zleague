package tournament

import (
	"context"
	"log"
	"time"
	"zleague/api/models"
	"github.com/robfig/cron"
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

func (t *Tournament) Insert(db *mongo.Database) {
	coll := db.Collection("tournament")

	_, err := coll.InsertOne(context.TODO(), t)
	if err != nil {
		log.Println(err)
	}
}

func (t *Tournament) UpdateInDB(db *mongo.Database) {
	coll := db.Collection("tournament")

	filter := bson.M{
		"starttime": t.StartTime,
	}

	update := bson.M{
		"$mod": {"teams": t.Teams},
	}

	err := coll.FindOneAndUpdate(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
}
