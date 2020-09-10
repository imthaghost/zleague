package tournament

import (
	"fmt"
	"zleague/backend_v2/models"
)

func createTeams(t map[string]TeamBasic) []models.Team {
	var allTeams []models.Team

	basicChan := make(chan TeamBasic, len(t))
	teamChan := make(chan models.Team, len(t))

	fmt.Println("start worker")
	for i := 0; i < 20; i++ {
		go teamWorker(basicChan, teamChan)
	}
	fmt.Println("20 workers started")
	for _, team := range t {
		fmt.Println(len(team.Teammates))
		basicChan <- team
	}
	close(basicChan)

	fmt.Println("All in worker")
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
