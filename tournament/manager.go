package tournament

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	cmap "github.com/orcaman/concurrent-map"
	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Manager is designed to manage multiple tournaments at once.
// It does this through a database, where it stores current tournaments and information about them.
// You can create new tournaments, delete tournaments, and more.
type Manager struct {
	Tournaments cmap.ConcurrentMap
	cron        *cron.Cron
	DB          *mongo.Database
}

// NewManager will create a new instance of tournament manager.
// By default, it will load tournaments that are currently in the database so that they can be interacted with.
func NewManager(db *mongo.Database) *Manager {
	// create a new tournament manager
	// m := &Manager{
	// 	Tournaments: map[string]Tournament{},
	// 	DB:          db,
	// }
	m := &Manager{
		Tournaments: cmap.New(),
		DB:          db,
	}

	// setup context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
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
		m.Tournaments.Set(tournament.ID, tournament)
		// m.Tournaments[tournament.ID] = tournament
	}

	return m
}

// Start will start a new tournament
func (t *Manager) Start() {
	// default to every 7 minutes
	schedule := "@every 7m"
	// create new cron instance for all our update loops
	c := cron.New()

	// start a new loop for every tournament
	log.Println("Cron Starting")
	for id, tourney := range t.Tournaments.Items() {
		log.Println("Starting Update Loop. Tournament ID: ", id)
		tournament := tourney.(Tournament)
		// start updating every x scheduled minutes
		err := c.AddFunc(schedule, updateLoop(t.DB, &tournament, t))
		if err != nil {
			log.Printf("Error adding update loop on start. ID: %s", tournament.ID)
		}
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
	teams := CreateTeams(start, end, csvData)
	newTournament := NewTournament(teams, id, start, end)

	err := newTournament.Insert(t.DB)
	if err != nil {
		log.Println("manager: error creating new tournament in db: ", err)
	}

	t.Tournaments.Set(newTournament.ID, newTournament)

	schedule := "@every 7m"
	// start updating every x scheduled minutes for the new tournament
	err = t.cron.AddFunc(schedule, updateLoop(t.DB, &newTournament, t))
	if err != nil {
		log.Printf("Error adding update loop on start. ID: %s", newTournament.ID)
	}

	return newTournament
}

func updateLoop(db *mongo.Database, t *Tournament, m *Manager) func() {
	return func() {
		// if time is before the time of the tournament, do nothing
		if time.Now().Before(t.StartTime) {
			log.Println("Tournament has not started yet... not updating..")
			return
		}

		log.Println("Updating Tournament. ID: ", t.ID)

		// log.Println("Current time " + time.Now().Format(time.RFC3339))
		// // we stop the cron job 30 minutes after the tournament endtime
		// if time.Now().After(tournament.EndTime.Add(time.Minute * time.Duration(30))) {
		//     // graceful kill
		//     log.Println("stopping cron")
		//     c.Stop()
		// }

		// Update all the teams
		t.Update()
		t.UpdateInDB(db)
		log.Println("Done Updating Tournament. ID: ", t.ID)

		// TODO: Update the tournament manager in memory with the updated tournament?
		tourney := Tournament{}
		tournament, err := tourney.GetTournament(db, t.ID)
		if err != nil {
			log.Println("manager: error getting tournament: ")
		}
		// update the tournament manager hashmap at once
		m.Tournaments.Set(t.ID, tournament)
	}
}

// GetTournament will get a tournament from memory
// GET /tournament/:id
// GET /tournament/:id/teams/:id
func (t *Manager) GetTournament(db *mongo.Database, id string) (Tournament, error) {
	tourney, exists := t.Tournaments.Get(id)
	if !exists {
		return Tournament{}, errors.New("manager: tournament does not exist")
	}

	// convert the interface to a tournament structure
	return tourney.(Tournament), nil
}

// AllTournaments will return all current active tournaments
func (t *Manager) AllTournaments() ([]Tournament, error) {
	// TODO: logic
	return []Tournament{}, nil
}
