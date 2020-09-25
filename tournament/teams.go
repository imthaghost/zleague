package tournament

import (
	"net/http"
	"zleague/api/models"
	"zleague/api/proxy"
)

// creates all the teams concurrently
func createTeams(t map[string]models.TeamBasic) []models.Team {
	// client
	c := proxy.NewNetClient()
	var allTeams []models.Team

	// instantiate two channels to pass the teams through
	basicChan := make(chan models.TeamBasic, len(t))
	teamChan := make(chan models.Team, len(t))

	// start 20 goroutines
	for i := 0; i < 50; i++ {
		go teamWorker(basicChan, teamChan, c)
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
	return allTeams
}

// CreateTeam instantiates a default Team
func createTeam(t models.TeamBasic, client *http.Client) models.Team {
	team := models.Team{
		Name:     t.Teamname,
		Players:  []models.Player{},
		Division: t.Division,
		Best:     models.Best{},
		Total:    models.Total{},
	}

	// for every player that is on the team, create a player object and add them to the players list
	for _, player := range t.Teammates {
		p := CreatePlayer(player, t.Teamname, client)
		team.Players = append(team.Players, p)
	}
	return team
}

// worker to concurrently create all the teams
func teamWorker(basic chan models.TeamBasic, team chan models.Team, client *http.Client) {
	// checks the channel for team objects and passes created teams into the team channel
	for t := range basic {
		team <- createTeam(t, client)
	}
}
