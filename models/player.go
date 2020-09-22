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
func (player *Player) getMatches(matches cod.MatchData, rules *Rules) {
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

	// gets run on first game
	if len(matches.Data.Matches) != 0 {
		player.LastMatch = matches.Data.Matches[0].Attributes.ID
	}
}
