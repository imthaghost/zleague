package tournament

import (
	"errors"
	"log"
	"net/http"
	"zleague/api/models"
	"zleague/api/proxy"

	cmap "github.com/orcaman/concurrent-map"
)

// creates all the teams concurrently
func createTeams(t map[string]models.TeamBasic) []models.Team {
	// client
	c := proxy.NewNetClient()
	var allTeams []models.Team

	// instantiate two channels to pass the teams through
	basicChan := make(chan models.TeamBasic, len(t))
	teamChan := make(chan models.Team, len(t))

	blockedTeams := cmap.New()
	// start 50 goroutines
	for i := 0; i < 50; i++ {
		go teamWorker(basicChan, teamChan, c, &blockedTeams)
	}
	// for every team in the map, add each to the channel
	for _, team := range t {
		basicChan <- team
	}
	// close the channel, not used again
	close(basicChan)

	// unload the teamChan into the allTeams array
	for i := 0; i < len(t); i++ {
		allTeams = append(allTeams, <-teamChan)
	}

	// display blocked teams
	for k, v := range blockedTeams.Items() {
		log.Printf("Blocked Team. TeamName: %s -- Player Name: %s\n", k, v)
	}

	// remove the empty teams for the ones that have blocked users or w/e
	allTeams = allTeams[:len(allTeams)-len(blockedTeams.Items())]

	return allTeams
}

// CreateTeam instantiates a default Team
func CreateTeam(t models.TeamBasic, client *http.Client, blocked *cmap.ConcurrentMap) (models.Team, error) {
	team := models.Team{
		Name:     t.Teamname,
		Players:  []models.Player{},
		Division: t.Division,
		Best:     models.Best{},
		Total:    models.Total{},
	}

	// for every player that is on the team, create a player object and add them to the players list
	for _, player := range t.Teammates {
		p, err := CreatePlayer(player, client)
		if err != nil {
			// return and do not create team, add team to blacklist and sepcify the user that failed them to be removed
			blocked.Set(t.Teamname, player)

			return models.Team{}, errors.New("did not create team. maybe invalid activison id")
		}

		team.Players = append(team.Players, p)
	}

	return team, nil
}

// worker to concurrently create all the teams
func teamWorker(basic chan models.TeamBasic, team chan models.Team, client *http.Client, blocked *cmap.ConcurrentMap) {
	// checks the channel for team objects and passes created teams into the team channel
	for tBasic := range basic {
		t, err := CreateTeam(tBasic, client, blocked)
		// if error, do not add the team
		if err != nil {
			log.Println("error when creating team. not adding to tournament and blocking...")
			team <- t
			continue
		}

		team <- t
	}
}
