package tournament

import (
	"log"
	"net/http"
	"sort"
	"time"
	"zleague/api/cod"
	"zleague/api/models"
)

// CreatePlayer creates a default Player instance
func CreatePlayer(username, teamname string, startTime, endTime time.Time, client *http.Client) models.Player {
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
		TStartTime: startTime,
		TEndTime:   endTime,
		Username:   username,
		Meta:       meta,
		Best:       best,
		Total:      total,
	}

	return player
}

// playerWorker is a goroutine worker to update each individual player concurrently
func playerWorker(playerChan chan *models.Player, client *http.Client, fin chan bool) {
	// checks the channel for players to update
	for player := range playerChan {
		// updates them and checks for errors
		all, err := cod.GetMatchData(player.Username, client)
		if err != nil {
			log.Println(err)
		}
		// add the matches from the updateStats function to the player
		updateStats(player, all)
		// finally update the stats of the player
		updatePlayer(player)
		// pass a value into the finalize channel, can be used to create a progress bar
		fin <- true
	}
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

// sortMatches sorts the matches and returns the n best matches
func sortMatches(matches []models.Match, n int) []models.Match {
	sort.Sort(models.ByScore(matches))
	// slice the array if more than n
	if len(matches) > n {
		matches = matches[:n]
	}
	return matches
}
