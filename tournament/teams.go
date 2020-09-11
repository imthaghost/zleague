package tournament

import (
	"zleague/api/models"
)

// creates all the teams concurrently
func createTeams(t map[string]TeamBasic) []models.Team {
	var allTeams []models.Team

	// instantiate two channels to pass the teams through
	basicChan := make(chan TeamBasic, len(t))
	teamChan := make(chan models.Team, len(t))

	// start 20 goroutines
	for i := 0; i < 20; i++ {
		go teamWorker(basicChan, teamChan)
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
func createTeam(t TeamBasic) models.Team {
	team := models.Team{
		Teamname:    t.Teamname,
		Kills:       0,
		Deaths:      0,
		Assists:     0,
		Headshots:   0,
		KD:          0.0,
		DamageDone:  0,
		GamesPlayed: 0,
		Wins:        0,
		Players:     []models.Player{},
		Division:    t.Division,
	}
	team.Total.TotalKills = 0
	team.Total.TotalAssists = 0
	team.Total.TotalDamage = 0
	team.Total.TotalDeaths = 0
	team.Total.TotalKD = 0
	team.Total.TotalHeadshots = 0
	team.Total.TotalWins = 0

	// for every player that is on the team, create a player object and add them to the players list
	for _, player := range t.Teammates {
		team.Players = append(team.Players, CreatePlayer(player, t.Teamname, t.Start, t.End))
	}
	return team
}

// worker to concurrently create all the teams
func teamWorker(basic chan TeamBasic, team chan models.Team) {
	// checks the channel for team objects and passes created teams into the team channel
	for t := range basic {
		team <- createTeam(t)
	}
}

// updates the team stats based off of the players stats
func updateTeam(team *models.Team) *models.Team {
	team.Kills = 0
	team.DamageDone = 0
	team.Deaths = 0
	team.Assists = 0
	team.Headshots = 0
	team.KD = float64(0)
	team.Wins = 0

	for i, player := range team.Players {
		team.Kills += player.Kills
		team.Deaths += player.Deaths
		team.Assists += player.Assists
		team.Headshots += player.Headshots
		team.KD = (float64(team.Kills) / float64(team.Deaths))
		team.DamageDone += player.DamageDone
		team.Wins = player.Wins
		team.TotalPoints = player.PlacementPoints
		team.PlacementPoints = player.PlacementPoints

		team.Total.TotalKills += player.Total.TotalKills
		team.Total.TotalDeaths += player.Total.TotalDeaths
		team.Total.TotalAssists += player.Total.TotalAssists
		team.Total.TotalHeadshots += player.Total.TotalHeadshots
		team.Total.TotalKD = (float64(team.Total.TotalKills) / float64(team.Total.TotalDeaths))
		team.Total.TotalDamage += player.Total.TotalDamage
		team.Total.TotalWins = player.Total.TotalWins
		team.Total.TotalScore = player.Total.TotalScore
		team.Players[i] = player
		team.GamesPlayed = player.GamesPlayed
	}
	// Scores
	team.TotalPoints += team.Kills
	team.Total.TotalScore += team.Total.TotalKills

	return team
}

// Update the teams and all of the players on the team.
func (t *Tournament) Update() {
	// instantiate 4 channels to use to pass the teams through,
	// one for the starting teams, updated teams, players and one finalize channel
	teamChan := make(chan *models.Team, 1000)
	updateChan := make(chan *models.Team, 1000)
	player := make(chan *models.Player, 1000)
	fin := make(chan bool, 1000)

	// instantiate 20 workers on each goroutine
	// 20 is the max amount of workers before rate limiting from the API
	for i := 0; i < 20; i++ {
		go updateWorker(teamChan, player)
		go playerWorker(player, fin)
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
