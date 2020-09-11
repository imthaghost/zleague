package tournament

import (
	"fmt"
	"log"
	"sort"
	"time"
	"zleague/backend_v2/cod"
	"zleague/backend_v2/models"
)

// CreatePlayer creates a default Player instance
func CreatePlayer(username, teamname string, startTime, endTime time.Time) models.Player {
	stats, err := cod.GetWarzoneStats(username)
	if err != nil {
		// TODO: Figure out what to do if player use
		log.Println(err)
		return models.Player{}
	}

	player := models.Player{
		TournamentStartTime: startTime,
		TournamentEndTime:   endTime,
		Username:            username,
		Teamname:            teamname,
		Avatar:              stats.Data.PlatformInfo.AvatarURL,
		Kills:               0,
		Deaths:              0,
		Headshots:           0,
		KD:                  0.0,
		PlacementPoints:     0,
		Assists:             0,
		DamageDone:          0,
		GamesPlayed:         0,
		LastMatch:           "",
		Matches:             []models.Match{},
		BestMatches:         []models.Match{},
	}
	player.Total.TotalKills = 0
	player.Total.TotalAssists = 0
	player.Total.TotalDamage = 0
	player.Total.TotalDeaths = 0
	player.Total.TotalKD = 0
	player.Total.TotalHeadshots = 0
	player.Level.Value = stats.Data.Segments[0].Stats.Level.Value
	player.Level.LevelIcon = stats.Data.Segments[0].Stats.Level.Metadata.IconURL

	return player
}

func playerWorker(playerChan chan *models.Player, fin chan bool) {
	for player := range playerChan {
		all, err := cod.GetWarzoneMatches(player.Username)
		if err != nil {
			log.Println(all)
		}

		player.Matches = updateStats(player, all)

		player.BestMatches = sortMatches(player.Matches, 4)

		updatePlayer(player)
		fin <- true
	}
}

func updatePlayer(player *models.Player) {
	player.Kills = 0
	player.Deaths = 0
	player.Assists = 0
	player.Headshots = 0
	player.KD = float64(0)
	player.DamageDone = 0
	player.PlacementPoints = 0

	for _, match := range player.BestMatches {
		player.Kills += match.Kills
		player.Deaths += match.Deaths
		player.Assists += match.Assists
		player.Headshots += match.Headshots
		player.KD = (float64(player.Kills) / float64(player.Deaths))
		player.DamageDone += match.DamageDone
		player.PlacementPoints += match.Score
	}
	fmt.Println(player.Total.TotalKills)
}

func sortMatches(matches []models.Match, num int) []models.Match {
	sort.Sort(models.ByScore(matches))
	if len(matches) > num {
		matches = matches[:num]
	}
	return matches
}
