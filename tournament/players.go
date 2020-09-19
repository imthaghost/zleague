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
func CreatePlayer(username, teamname string, startTime, endTime time.Time) models.Player {
	stats, err := cod.GetWarzoneStats(username)
	if err != nil {
		// TODO: Figure out what to do if player username doesnt exist
		log.Println(err)
		return models.Player{}
	}

	// instantiate a player model to store the data in.
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

// goroutine worker to update each individual player concurrently
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

		// sort and add the best matches into the bestMatches array
		player.BestMatches = sortMatches(player.Matches, 4)
		// finally update the stats of the player
		updatePlayer(player)
		// pass a value into the finalize channel, can be used to create a progress bar
		fin <- true
	}
}

// updates the stats of the player based off of the matches they have stored
func updatePlayer(player *models.Player) {
	// resets the stats of the player to zero
	player.Kills = 0
	player.Deaths = 0
	player.Assists = 0
	player.Headshots = 0
	player.KD = float64(0)
	player.DamageDone = 0
	player.PlacementPoints = 0

	// iterate over all of the matches and update the stats
	for _, match := range player.BestMatches {
		player.Kills += match.Kills
		player.Deaths += match.Deaths
		player.Assists += match.Assists
		player.Headshots += match.Headshots
		player.KD = (float64(player.Kills) / float64(player.Deaths))
		player.DamageDone += match.DamageDone
		player.PlacementPoints += match.Score
	}
}

// sorts the matches and returns the n best matches
func sortMatches(matches []models.Match, n int) []models.Match {
	sort.Sort(models.ByScore(matches))
	// slice the array if more than n
	if len(matches) > n {
		matches = matches[:n]
	}
	return matches
}
