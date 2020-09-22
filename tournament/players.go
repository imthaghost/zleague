package tournament

import (
	"log"
	"net/http"
	"zleague/api/cod"
	"zleague/api/models"
)

// CreatePlayer creates a default Player instance
func CreatePlayer(username, teamname string, client *http.Client) models.Player {
	stats, err := cod.GetStatData(username, client)
	if err != nil {
		log.Println(err)
		return models.Player{}
	}

	best := models.Best{}
	total := models.Total{}
	meta := models.Meta{
		Avatar:    stats.Data.PlatformInfo.AvatarURL,
		Level:     stats.Data.Segments[0].Stats.Level.Value,
		LevelIcon: stats.Data.Segments[0].Stats.Level.Metadata.IconURL,
	}

	player := models.Player{
		Username: username,
		Meta:     meta,
		Best:     best,
		Total:    total,
	}

	return player
}

// updates the stats of the player based off of the matches they have stored
func updatePlayer(player *models.Player) {
	// resets the stats of the player to zero
	best := models.Best{}

	// iterate over all of the matches and update the stats
	for _, match := range player.Best.Games {
		// increment player wins
		if match.Placement == 1 {
			best.Wins++
		}

		best.Kills += match.Kills
		best.Deaths += match.Deaths
		best.Headshots += match.Headshots
		best.KD = (float64(best.Kills) / float64(best.Deaths))
		best.DamageDone += match.DamageDone
		best.PlacementPoints += match.Score
	}
	best.Games = player.Best.Games
	player.Best = best
}
