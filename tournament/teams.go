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
	for i := 0; i < 20; i++ {
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
		p := CreatePlayer(player, t.Teamname, t.Start, t.End, client)
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

// updates the team stats based off of the players stats
func updateTeam(team *models.Team) *models.Team {
	best := models.Best{}
	total := models.Total{}

	for i, player := range team.Players {
		best.Kills += player.Best.Kills
		best.Deaths += player.Best.Deaths
		best.Headshots += player.Best.Headshots

		// update kd cause its special and we dont like to divide by 0
		if best.Deaths == 0 {
			best.KD = float64(best.Kills)
		} else {
			best.KD = (float64(best.Kills) / float64(best.Deaths))
		}

		best.DamageDone += player.Best.DamageDone
		best.Wins = player.Best.Wins
		best.CombinedPoints = player.Best.PlacementPoints
		best.PlacementPoints = player.Best.PlacementPoints

		total.Kills += player.Total.Kills
		total.Deaths += player.Total.Deaths
		total.Headshots += player.Total.Headshots
		if total.Deaths == 0 {
			total.KD = float64(total.Kills)
		} else {
			total.KD = (float64(total.Kills) / float64(total.Deaths))
		}
		total.DamageDone += player.Total.DamageDone
		total.Wins = player.Total.Wins
		total.CombinedPoints = player.Total.CombinedPoints
		team.Players[i] = player
		total.GamesPlayed = player.Total.GamesPlayed
	}

	// Scores
	best.CombinedPoints += best.Kills
	total.CombinedPoints += total.Kills

	team.Best = best
	team.Total = total

	return team
}

// updateWorker goroutine handles updating all of the players on the team
func updateWorker(teamChan chan *models.Team, playerChan chan *models.Player) {
	// check the channel for any teams passed in
	for t := range teamChan {
		for i := range t.Players {
			playerChan <- &t.Players[i]
		}
	}
}

// updateTeamStatsWorker updates the stats on the team based off of the updated players
func updateTeamStatsWorker(teamChan chan *models.Team, fin chan bool) {
	// check the channel for any teams passed in
	for team := range teamChan {
		updateTeam(team)
		fin <- true
	}
}
