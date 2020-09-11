package tournament

import (
	"fmt"
	"zleague/backend_v2/models"
)

func createTeams(t map[string]TeamBasic) []models.Team {
	var allTeams []models.Team

	basicChan := make(chan TeamBasic, len(t))
	teamChan := make(chan models.Team, len(t))

	for i := 0; i < 20; i++ {
		go teamWorker(basicChan, teamChan)
	}
	for _, team := range t {
		basicChan <- team
	}
	close(basicChan)

	for i := 0; i < len(t); i++ {
		allTeams = append(allTeams, <-teamChan)
	}
	fmt.Println("Created Teams")
	return allTeams
}

// CreateTeam creates a default Team instance
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

	for _, player := range t.Teammates {
		team.Players = append(team.Players, CreatePlayer(player, t.Teamname, t.Start, t.End))
	}
	return team
}

func teamWorker(basic chan TeamBasic, team chan models.Team) {
	for t := range basic {
		team <- createTeam(t)
	}
}

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

// UpdateTeam does
func (t *Tournament) UpdateTeam() {
	teamChan := make(chan *models.Team, 1000)
	updateChan := make(chan *models.Team, 1000)
	player := make(chan *models.Player, 1000)
	fin := make(chan bool, 1000)

	fmt.Println("start worker")
	for i := 0; i < 20; i++ {
		go updateWorker(teamChan, player)
		go playerWorker(player, fin)
		go updateTeamStatsWorker(updateChan, fin)
	}
	fmt.Println("20 workers started")

	for i := range t.Teams {
		teamChan <- &t.Teams[i]
	}

	for i := 0; i < 1; i++ {
		<-fin
	}

	for i := range t.Teams {
		updateChan <- &t.Teams[i]
	}

	for i := 0; i < 1; i++ {
		<-fin
	}
}

func updateWorker(teamChan chan *models.Team, playerChan chan *models.Player) {
	for t := range teamChan {
		for i := range t.Players {
			playerChan <- &t.Players[i]
		}
	}
}

func updateTeamStatsWorker(teamChan chan *models.Team, fin chan bool) {
	for team := range teamChan {
		updateTeam(team)
		fin <- true
	}
}
