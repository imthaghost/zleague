package tournament

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Manager is designed to manage multiple tournaments at once.
// It does this through a database, where it stores current tournaments and information about them.
// You can create new tournaments, delete tournaments, and more.
type Manager struct {
	Tournaments map[string]Tournament // represents all current tournaments
	client      http.Client
	cron        *cron.Cron
	DB          *mongo.Database
}

// NewManager will create a new instance of tournament manager.
// By default, it will load tournaments that are currently in the database so that they can be interacted with.
func NewManager(db *mongo.Database) *Manager {
	m := &Manager{
		Tournaments: map[string]Tournament{},
		DB:          db,
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
		m.Tournaments[tournament.ID] = tournament
	}

	return m
}

// Start will start a new tournament
func (t *Manager) Start() {
	// default to every 10 minutes
	schedule := "@every 1m"
	// create new cron instance
	c := cron.New()

	// start a new loop for every tournament
	log.Println("Cron Starting")
	for id, tournament := range t.Tournaments {
		log.Println("Starting Update loop for tournament id: ", id)
		// call function every 10 minutes
		c.AddFunc(schedule, func() {
			// if time is before the time of the tournament, do nothing
			if time.Now().Before(tournament.StartTime) {
				log.Println("Tournament has not started yet... not updating..")
				return
			}

			// log.Println("Current time " + time.Now().Format(time.RFC3339))
			// // we stop the cron job 30 minutes after the tournament endtime
			// if time.Now().After(tournament.EndTime.Add(time.Minute * time.Duration(30))) {
			//     // graceful kill
			//     log.Println("stopping cron")
			//     c.Stop()
			// }

			// Update all the teams
			tournament.Update()
			tournament.UpdateInDB(t.DB)
		})
	}

	// start the jobs
	c.Start()

	// bind to tournament manager so we can start more crons and what not
	t.cron = c
}

// NewTournament is designed to create a new tournament, and then save it to the struct and the database and return it
func (t *Manager) NewTournament(start, end time.Time, id string, csvData io.Reader) Tournament {
	// create a new tournament
	// TODO: Start the cron job for this tournament because it wont be started from the "start"
	teams := Create(start, end, csvData)
	newTournament := NewTournament(teams, id, start, end)

	err := newTournament.Insert(t.DB)
	if err != nil {
		log.Println("manager: error creating new tournament in db: ", err)
	}

	t.Tournaments[newTournament.ID] = newTournament

	schedule := "@every 1m"
	t.cron.AddFunc(schedule, func() {
		// log.Println("Current time " + time.Now().Format(time.RFC3339))
		// // we stop the cron job 30 minutes after the tournament endtime
		// if time.Now().After(tournament.EndTime.Add(time.Minute * time.Duration(30))) {
		//     // graceful kill
		//     log.Println("stopping cron")
		//     c.Stop()
		// }
		// if time is before the time of the tournament, do nothing
		if time.Now().Before(newTournament.StartTime) {
			log.Println("Tournament has not started yet... not updating...")
			return
		}
		// Update all the teams
		log.Println("updating due to create")
		newTournament.Update()
		newTournament.UpdateInDB(t.DB)
	})

	return newTournament
}

// GetTournament will get a tournament from memory or w/e
// GET /tournament/:id
// GET /tournament/:id/teams/:id
func (t *Manager) GetTournament(db *mongo.Database, id string) Tournament {
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

// AllTournaments will return all current active tournaments
func (t *Manager) AllTournaments() []Tournament {
	// TODO: logic
	return []Tournament{}
}
