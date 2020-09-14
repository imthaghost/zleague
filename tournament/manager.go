package tournament

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TournamentManager is designed to manage multiple tournaments at once.
// It does this through a database, where it stores current tournaments and information about them.
// You can create new tournaments, delete tournaments, and more.
type TournamentManager struct {
	tournaments map[string]Tournament // represents all current tournaments
	client      http.Client
	cron        *cron.Cron
}

// NewTournamentManager will create a new instance of tournament manager.
// By default, it will load tournaments that are currently in the database so that they can be interacted with.
func NewTournamentManager(db *mongo.Database) *TournamentManager {
	m := &TournamentManager{
		tournaments: map[string]Tournament{},
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	// in here, we want to load the tournaments that already exist
	cursor, err := db.Collection("tournaments").Find(context.TODO(), bson.D{})
	if err != nil {
		defer cursor.Close(ctx)
		log.Println("manager: error finding tournaments: ", err)
	}

	// loop through the different tournaments
	for cursor.Next(ctx) {
		var tournament Tournament
		err := cursor.Decode(&tournament)

		if err != nil {
			log.Println("manager: error decoding tournament: ", err)
		}

		// add the tournament to the map
		m.tournaments[tournament.ID] = tournament
	}

	return m
}

func (t *TournamentManager) Start() {
	// START CRONS
	// default to every 10 minutes
	schedule := "@every 1m"
	// create new cron instance
	c := cron.New()

	// start a new loop for every tournament
	log.Println("Starting cron:")
	for id, tournament := range t.tournaments {
		log.Println("Starting Update loop for tournament id: ", id)
		// call function every 10 minutes
		c.AddFunc(schedule, func() {
			// log.Println("Current time " + time.Now().Format(time.RFC3339))
			// // we stop the cron job 30 minutes after the tournament endtime
			// if time.Now().After(tournament.EndTime.Add(time.Minute * time.Duration(30))) {
			//     // graceful kill
			//     log.Println("stopping cron")
			//     c.Stop()
			// }

			// Update all the teams
			log.Println("updating due to start")
			tournament.Update()
		})
	}

	// start the jobs
	c.Start()

	// bind to tournament manager so we can start more crons and what not
	t.cron = c
}

// NewTourament is designed to create a new tournament, and then save it to the struct and the database and return it
func (t *TournamentManager) NewTournament(db *mongo.Database, start, end time.Time, id string) Tournament {
	// create a new tournament
	// TODO: Start the cron job for this tournament because it wont be started from the "start"
	teams := Create(start, end)
	newTournament := NewTournament(teams, id, start, end)

	_, err := db.Collection("tournaments").InsertOne(context.TODO(), newTournament)

	if err != nil {
		log.Println("manager: error creating new tournament in db: ", err)
	}

	t.tournaments[newTournament.ID] = newTournament

	schedule := "@every 1m"
	t.cron.AddFunc(schedule, func() {
		// log.Println("Current time " + time.Now().Format(time.RFC3339))
		// // we stop the cron job 30 minutes after the tournament endtime
		// if time.Now().After(tournament.EndTime.Add(time.Minute * time.Duration(30))) {
		//     // graceful kill
		//     log.Println("stopping cron")
		//     c.Stop()
		// }

		// Update all the teams
		log.Println("updating due to create")
		newTournament.Update()
	})

	return newTournament
}

// GetTournament will get a tournament from memory or w/e
// GET /tournament/:id
// GET /tournament/:id/teams/:id
func (t *TournamentManager) GetTournament(db *mongo.Database, id string) Tournament {
	// TODO: check to see if tournament is im memory before pulling from the db (or just update the map)
	var tournament Tournament
	coll := db.Collection("tournaments")

	err := coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "id", Value: id}}).Decode(&tournament)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return tournament
		}
		log.Println("manager: error getting tournament: ", err)
	}

	return tournament
}

func (t *TournamentManager) AllTournaments() []Tournament {
	// TODO: logic
	return []Tournament{}
}
