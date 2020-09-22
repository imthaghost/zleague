package tournament

import (
	"zleague/api/cod"
	"zleague/api/models"
)

// updates the stats of a given player, takes a player and a list of matches as an argument

// ONLY ADD MATCHES TO TEAM/PLAYER
func updateStats(player *models.Player, matches cod.MatchData) {
	// iterate over the matches
	for _, match := range matches.Data.Matches {
		// checks to make sure the match is during the tournament times and is an allowed type of match
		if match.Attributes.ID == player.LastMatch {
			break
		} else if match.Metadata.Timestamp.Before(player.TStartTime) {
			break
		} else if match.Metadata.Timestamp.After(player.TEndTime) {
			continue
		} else if match.Attributes.ModeID != "br_brtrios" {
			continue
		}

		var kd float64
		if match.Segments[0].Stats.Deaths.Value == 0 {
			kd = float64(match.Segments[0].Stats.Kills.Value)
		} else {
			kd = match.Segments[0].Stats.Kills.Value / match.Segments[0].Stats.Deaths.Value
		}

		// create a new match structure and store the data from the API
		newMatch := models.Match{
			ID:          match.Attributes.ID,
			Mode:        match.Metadata.ModeName,
			StartTime:   match.Metadata.Timestamp,
			Kills:       int(match.Segments[0].Stats.Kills.Value),
			Deaths:      int(match.Segments[0].Stats.Deaths.Value),
			Assists:     int(match.Segments[0].Stats.Assists.Value),
			Headshots:   int(match.Segments[0].Stats.Headshots.Value),
			KD:          kd,
			TimePlayed:  match.Segments[0].Stats.TeamSurvivalTime.DisplayValue,
			Placement:   match.Segments[0].Stats.Placement.Value,
			DamageDone:  int(match.Segments[0].Stats.DamageDone.Value),
			DamageTaken: int(match.Segments[0].Stats.DamageTaken.Value),
			Score:       Scoreboard[match.Segments[0].Stats.Placement.Value],
		}

		// append the matches into the player reference and into a newMatches array
		player.Matches = append(player.Matches, newMatch)

		// add stats to the players total
		player.Total.Kills += int(match.Segments[0].Stats.Kills.Value)
		player.Total.Assists += int(match.Segments[0].Stats.Assists.Value)
		player.Total.DamageDone += int(match.Segments[0].Stats.DamageDone.Value)
		player.Total.DamageTaken += int(match.Segments[0].Stats.DamageTaken.Value)
		player.Total.Deaths += int(match.Segments[0].Stats.Deaths.Value)

		if player.Total.Deaths == 0 {
			player.Total.KD = float64(player.Total.Kills)
		} else {
			player.Total.KD = (float64(player.Total.Kills) / float64(player.Total.Deaths))
		}

		player.Total.Headshots += int(match.Segments[0].Stats.Headshots.Value)
		player.Total.CombinedPoints += newMatch.Score
		player.Total.GamesPlayed++

		// check if player won the game
		if newMatch.Placement == 1 {
			player.Total.Wins++
		}
	}

	// gets run on first game
	if len(matches.Data.Matches) != 0 {
		player.LastMatch = matches.Data.Matches[0].Attributes.ID
	}
}

// not done, but will allow to check past tournaments
// func updateAll(player *models.Player, matches []cod.MatchData) []models.Match {
// 	var newMatches []models.Match

// 	for j, m := range matches {
// 		for i, match := range m.Data.Matches {
// 			if match.Metadata.Timestamp.Before(player.TournamentStartTime) {
// 				if i != 0 {
// 					player.LastMatch = m.Data.Matches[j].Attributes.ID
// 				}
// 				break
// 			} else if match.Metadata.Timestamp.After(player.TournamentEndTime) {
// 				continue
// 			} else if match.Attributes.ID == player.LastMatch {
// 				player.LastMatch = m.Data.Matches[j].Attributes.ID
// 				break
// 			} else if match.Attributes.ModeID != "br_brtrios" {
// 				continue
// 			}

// 			newMatch := models.Match{
// 				ID:         match.Attributes.ID,
// 				Mode:       match.Metadata.ModeName,
// 				StartTime:  match.Metadata.Timestamp,
// 				Kills:      int(match.Segments[0].Stats.Kills.Value),
// 				Deaths:     int(match.Segments[0].Stats.Deaths.Value),
// 				Assists:    int(match.Segments[0].Stats.Assists.Value),
// 				Headshots:  int(match.Segments[0].Stats.Headshots.Value),
// 				KD:         (match.Segments[0].Stats.Kills.Value / match.Segments[0].Stats.Deaths.Value),
// 				TimePlayed: match.Segments[0].Stats.TeamSurvivalTime.DisplayValue,
// 				Placement:  match.Segments[0].Stats.Placement.Value,
// 				DamageDone: int(match.Segments[0].Stats.DamageDone.Value),
// 				Score:      Scoreboard[match.Segments[0].Stats.Placement.Value],
// 			}

// 			player.Matches = append(player.Matches, newMatch)
// 			newMatches = append(newMatches, newMatch)

// 			player.Total.TotalKills += int(match.Segments[0].Stats.Kills.Value)
// 			player.Total.TotalAssists += int(match.Segments[0].Stats.Assists.Value)
// 			player.Total.TotalDamage += int(match.Segments[0].Stats.DamageDone.Value)
// 			player.Total.TotalDeaths += int(match.Segments[0].Stats.Deaths.Value)
// 			player.Total.TotalKD = (float64(player.Total.TotalKills) / float64(player.Total.TotalDeaths))
// 			player.Total.TotalHeadshots += int(match.Segments[0].Stats.Headshots.Value)
// 			player.Total.TotalScore += newMatch.Score
// 			player.GamesPlayed++

// 			if newMatch.Placement == 1 {
// 				player.Total.TotalWins++
// 			}
// 		}
// 	}
// 	return newMatches
// }
