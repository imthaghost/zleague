package tournament

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
