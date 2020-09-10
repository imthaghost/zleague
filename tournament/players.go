package tournament

import (
	"log"
	"time"
	"zleague/backend_v2/cod"
	"zleague/backend_v2/models"
)

// CreatePlayer creates a default Player instance
func CreatePlayer(username, teamname string, startTime, endTime time.Time) models.Player {
	stats, err := cod.GetWarzoneStats(username)
	if err != nil {
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
		Score:               0,
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
