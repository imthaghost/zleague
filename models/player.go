package models

import (
	"zleague/api/cod"
)

// Player represents a single player in the tournament
type Player struct {
	Username  string `json:"username"`
	Teamname  string `json:"team_name"`
	LastMatch string `json:"-"`
	Total     Total  `json:"total"` // this struct holds the total games played
	Best      Best   `json:"best"`  // this struct holds the best x games as determined when a tournament is created
	Meta      Meta   `json:"meta"`  // meta represents random meta data like leven and icon
}

// Meta represents meta data about the player
type Meta struct {
	Level     int    `json:"level"`
	LevelIcon string `json:"level_icon"`
	Avatar    string `json:"avatar"`
}

// getMatches will update all the stats an individual player has played
func (player *Player) getMatches(matches cod.MatchData, rules *Rules) *Player {
	// iterate over the matches
	for _, match := range matches.Data.Matches {
		// checks to make sure the match is during the tournament times and is an allowed type of match
		if match.Attributes.ID == player.LastMatch {
			break
		} else if match.Metadata.Timestamp.Before(rules.StartTime) {
			break
		} else if match.Metadata.Timestamp.After(rules.EndTime) {
			continue
		} else if match.Attributes.ModeID != rules.GameMode {
			continue
		}

		var kd float64
		if match.Segments[0].Stats.Deaths.Value == 0 {
			kd = float64(match.Segments[0].Stats.Kills.Value)
		} else {
			kd = match.Segments[0].Stats.Kills.Value / match.Segments[0].Stats.Deaths.Value
		}

		// create a new match structure and store the data from the API
		newMatch := Match{
			ID:          match.Attributes.ID,
			Mode:        match.Metadata.ModeName,
			StartTime:   match.Metadata.Timestamp,
			Kills:       int(match.Segments[0].Stats.Kills.Value),
			Deaths:      int(match.Segments[0].Stats.Deaths.Value),
			Headshots:   int(match.Segments[0].Stats.Headshots.Value),
			KD:          kd,
			TimePlayed:  match.Segments[0].Stats.TeamSurvivalTime.DisplayValue,
			Placement:   match.Segments[0].Stats.Placement.Value,
			DamageDone:  int(match.Segments[0].Stats.DamageDone.Value),
			DamageTaken: int(match.Segments[0].Stats.DamageTaken.Value),
			Score:       Scoreboard[match.Segments[0].Stats.Placement.Value],
			Checked:     false,
		}

		// append the matches into the player reference and into a newMatches array
		player.Total.Games = append(player.Total.Games, newMatch)
	}

	// checks if their are matches that were checked,
	// if so, makes the players last checked match equal to the first match returned
	if len(matches.Data.Matches) != 0 {
		player.LastMatch = matches.Data.Matches[0].Attributes.ID
	}
	// return a reference to the player
	return player
}

// updates the best stats on a player helper method
func (player *Player) updateBest() {
	// instantiates an empty best struct
	best := Best{}
	// for every game in the players best games, update the stats
	for _, match := range player.Best.Games {
		if match.Placement == 1 {
			best.Wins++
		}

		best.Kills += match.Kills
		best.Deaths += match.Deaths
		best.Headshots += match.Headshots
		best.KD = (float64(best.Kills) / float64(best.Deaths))
		best.DamageDone += match.DamageDone
		best.DamageTaken += match.DamageTaken
		best.PlacementPoints += match.Score
	}
	best.CombinedPoints = best.Kills + best.PlacementPoints
	best.Games = player.Best.Games
	// reassign players best struct to the new best struct
	player.Best = best
}

// updateMatches will update all the matches on a specific player in the seenMatches histogran
func (player *Player) updateMatches(seenMatches *map[string]Match) {
	// checks all the games on the player
	for j, m := range player.Total.Games {
		// if the match has been checked before, continue
		if m.Checked == true {
			continue
		}
		player.Total.Games[j].Checked = true
		// update the histogram with all the matches the player has played
		match, exists := (*seenMatches)[m.ID]
		// if the match exists in the map, update the stats of the match to reflect the teams total score
		if exists {
			match.Seen++
			match.Kills += m.Kills
			match.Deaths += m.Deaths
			match.Headshots += m.Headshots
			match.DamageDone += m.DamageDone
			match.DamageTaken += m.DamageTaken
			match.Score = m.Score
			match.Checked = true

			// catch divide by 0 error
			if match.Deaths == 0 {
				match.KD = float64(match.Kills)
			} else {
				match.KD = float64(match.Kills) / float64(match.Deaths)
			}
			(*seenMatches)[m.ID] = match
		} else {
			// not seen we insert the match into the map
			m.Seen++
			m.Checked = true
			(*seenMatches)[m.ID] = m
		}
	}
}

// updateStats updates all the stats on a player
func (player *Player) updateStats(seenMatches *map[string]Match, rules Rules) {
	// instantiate an empty variable to store the number of matches to be deleted
	var n int
	for j, m := range player.Total.Games {
		match, exists := (*seenMatches)[m.ID]
		if exists && match.Seen == rules.TeamSize {
			player.Total.Kills += m.Kills
			player.Total.DamageDone += m.DamageDone
			player.Total.DamageTaken += m.DamageTaken
			player.Total.Deaths += m.Deaths
			player.Total.PlacementPoints += m.Score

			if player.Total.Deaths == 0 {
				player.Total.KD = float64(player.Total.Kills)
			} else {
				player.Total.KD = (float64(player.Total.Kills) / float64(player.Total.Deaths))
			}

			player.Total.Headshots += m.Headshots
			player.Total.GamesPlayed++

			// check if player won the game
			if m.Placement == 1 {
				player.Total.Wins++
			}
		} else if match.Seen != rules.BestGamesNum {
			// if the match hasn't been seen the appropriate number of times in the histogram
			// swap the match with whatever match at the 0 + n index
			player.Total.Games[j], player.Total.Games[0+n] = player.Total.Games[0+n], player.Total.Games[j]
			// increment n
			n++
		}
	}
	// slice the players total games at n
	player.Total.Games = player.Total.Games[n:]
	player.Total.CombinedPoints = player.Total.Kills + player.Total.PlacementPoints
	// sort the matches and return a slice of the best matches according to the rules
	player.Best.Games = sortMatches(player.Total.Games, rules.BestGamesNum)

	// update the players best stats
	player.updateBest()
}
