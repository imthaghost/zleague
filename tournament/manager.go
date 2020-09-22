package tournament

import (
	"context"
	"errors"
	"io"
	"log"
	"sort"
	"time"
	"zleague/api/models"
	"zleague/api/proxy"

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
		var tournament models.Tournament
		err := cursor.Decode(&tournament)

		if err != nil {
			log.Println("manager: error decoding tournament: ", err)
		}

		// add the tournament to the map
		m.Tournaments.Set(tournament.ID, tournament)
	}

	return m
}

// Start will start a new tournament
func (m *Manager) Start() {
	// default to every 3 minutes
	schedule := "@every 3m"
	// create new cron instance for all our update loops
	c := cron.New()

	// start a new loop for every tournament
	log.Println("Cron Starting")
	for id, tourney := range m.Tournaments.Items() {
		log.Println("Starting Update Loop. Tournament ID: ", id)
		tournament := tourney.(models.Tournament)
		// start updating every x scheduled minutes
		err := c.AddFunc(schedule, updateLoop(m.DB, &tournament, m))
		if err != nil {
			log.Printf("Error adding update loop on start. ID: %s", tournament.ID)
		}
	}

	// start the jobs
	c.Start()

	// bind to tournament manager so we can start more crons and what not
	m.cron = c
}

// NewTournament is designed to create a new tournament, and then save it to the struct and the database and return it
func (m *Manager) NewTournament(start, end time.Time, id string, csvData io.Reader) models.Tournament {
	// create a new tournament
	// TODO: Start the cron job for this tournament because it wont be started from the "start"
	teams := CreateTeams(start, end, csvData)

	// TODO: create rules in route
	rules := models.Rules{
		StartTime:    start,
		EndTime:      end,
		BestGamesNum: 4,
		GameMode:     "br_brtrios",
	}

	newTournament := NewTournament(id, rules, teams)

	err := newTournament.Insert(m.DB)
	if err != nil {
		log.Println("manager: error creating new tournament in db: ", err)
	}

	m.Tournaments.Set(newTournament.ID, newTournament)

	schedule := "@every 3m"
	// start updating every x scheduled minutes for the new tournament
	err = m.cron.AddFunc(schedule, updateLoop(m.DB, &newTournament, m))
	if err != nil {
		log.Printf("Error adding update loop on start. ID: %s", newTournament.ID)
	}

	return newTournament
}

func updateLoop(db *mongo.Database, t *models.Tournament, m *Manager) func() {
	return func() {
		// if time is before the time of the tournament, do nothing
		if time.Now().Before(t.Rules.StartTime) {
			log.Println("Tournament has not started yet... not updating..")
			return
		}

		// we stop the cron job 30 minutes after the tournament endtime
		if time.Now().After(t.Rules.EndTime.Add(time.Minute * time.Duration(30))) {
			// graceful kill
			log.Println("Not running update. Tournament updated.")
			return
		}

		log.Println("Updating Tournament. ID: ", t.ID)

		// Update all the teams
		m.Update(t)
		t.UpdateInDB(db)
		log.Println("Done Updating Tournament. ID: ", t.ID)

		tourney := models.Tournament{}
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
func (m *Manager) GetTournament(id string) (models.Tournament, error) {
	tourney, exists := m.Tournaments.Get(id)
	if !exists {
		return models.Tournament{}, errors.New("manager: tournament does not exist")
	}

	// convert the interface to a tournament structure
	return tourney.(models.Tournament), nil
}

// AllTournaments will return all current active tournaments
func (m *Manager) AllTournaments() ([]models.Tournament, error) {
	// TODO: logic
	return []models.Tournament{}, nil
}

// GetTeam will return a single team that is within a tournament cache
func (m *Manager) GetTeam(id, name string) (models.Team, error) {
	t, exists := m.Tournaments.Get(id)
	if !exists {
		return models.Team{}, errors.New("tournament not found")
	}

	tournament := t.(models.Tournament)

	for _, team := range tournament.Teams {
		// if we find the team with the given name
		if team.Name == name {
			return team, nil
		}
	}

	// we did not find the team
	return models.Team{}, errors.New("team with that name not found within this tournament")
}

// GetTeams will return the teams that are in the cache
func (m *Manager) GetTeams(id string) ([]models.Team, error) {
	t, exists := m.Tournaments.Get(id)
	if !exists {
		return []models.Team{}, errors.New("manager: tournament does not exist")
	}

	tournament := t.(models.Tournament)

	return tournament.Teams, nil
}

// GetTeamsByDivision will return the teams in a given division
func (m *Manager) GetTeamsByDivision(id, div string) ([]models.Team, error) {
	t, exists := m.Tournaments.Get(id)
	if !exists {
		return []models.Team{}, errors.New("manager: tournament does not exist")
	}

	tournament := t.(models.Tournament)

	var teams []models.Team
	for _, team := range tournament.Teams {
		if team.Division == div {
			teams = append(teams, team)
		}
	}

	// handle incorrect division id
	if len(teams) == 0 {
		return teams, errors.New("no teams found for that divison / division was incorrect")
	}

	return teams, nil
}

// Update will update the given tournament
func (m *Manager) Update(t *models.Tournament) {
	// instantiate 4 channels to use to pass the teams through,
	// one for the starting teams, updated teams, players and one finalize channel
	teamChan := make(chan *models.Team, len(t.Teams)*2)
	updateChan := make(chan *models.Team, len(t.Teams)*2)
	player := make(chan *models.Player, len(t.Teams)*4)
	fin := make(chan bool, len(t.Teams)*4)
	client := proxy.NewNetClient() // sync http client

	// instantiate 30 workers on each goroutine
	// 30 is the max amount of workers before rate limiting from the API
	for i := 0; i < 30; i++ {
		go updateWorker(teamChan, player)
		go playerWorker(player, client, fin)
		go updateTeamStatsWorker(updateChan, fin)
	}

	//  for each team in the tournament, pass them through the channel
	for i := range t.Teams {
		teamChan <- &t.Teams[i]
	}

	// unload the finalize channel to know when the first channel finishes
	// iterate for the total number of players in the tournament
	for i := 0; i < (len(t.Teams) * 3); i++ {
		<-fin
	}

	// go through the teams again and pass them through the update channel
	// must be done after the finalize to make sure that the teams have been completely updated
	for i := range t.Teams {
		updateChan <- &t.Teams[i]
	}

	// unload the finalize channel one last time.
	// iterate for the total number of teams in the tournament
	for i := 0; i < len(t.Teams); i++ {
		<-fin
	}
	// Sort the teams by the number of points they have
	sort.Sort(models.ByPoints(t.Teams))
}
