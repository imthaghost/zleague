package models

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"zleague/api/cod"
)

// Team represents a single team in the tournament
type Team struct {
	Name     string   `json:"name"`
	Division string   `json:"division"`
	Best     Best     `json:"best"`  // this struct holds the best x games as determined when a tournament is created
	Total    Total    `json:"total"` // this struct hold the totals for ALL the games
	Players  []Player `json:"players"`
}

// TeamBasic holds a simple struct of what a team consists of.
type TeamBasic struct {
	Teamname  string    `json:"team_name"`
	Start     time.Time `json:"start_time"`
	End       time.Time `json:"end_time"`
	Division  string    `json:"division"`
	Teammates []string  `json:"teammates"`
}

// ByPoints allows us to sort all the teams
type ByPoints []Team

func (a ByPoints) Len() int           { return len(a) }
func (a ByPoints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPoints) Less(i, j int) bool { return a[i].Best.CombinedPoints > a[j].Best.CombinedPoints }

// Update will do all the update logic for a single team
func (t *Team) Update(client *http.Client, rules Rules) {
	// "histogram" for the matches that we have seen that we can then filter
	seenMatches := map[string]Match{}
	fmt.Println("starting update on ", t.Name)
	// go through all players on the team and update their "all matches" on the player model
	for i := range t.Players {
		// get the player info from warzone cod package
		matches, err := cod.GetMatchData(t.Players[i].Username, client)
		if err != nil {
			log.Println(err, "there was an error!")
		}
		// convert matches to match struct and store on the player
		t.Players[i].getMatches(matches, &rules)
		// working here
		// shit happens
		for j, m := range t.Players[i].Total.Games {
			t.Players[i].Total.Games[j].Checked = true
			// update the histogram with all the matches the player has played
			match, exists := seenMatches[m.ID]
			// if the match exists in the map, update the stats of the match to reflect the teams total score
			if exists {
				match.Seen++
				match.Kills += m.Kills
				match.Deaths += m.Deaths
				match.Headshots += m.Headshots
				match.DamageDone += m.DamageDone
				match.DamageTaken += m.DamageTaken
				match.Score = m.Score

				// catch divide by 0 error
				if match.Deaths == 0 {
					match.KD = float64(match.Kills)
				} else {
					match.KD = float64(match.Kills) / float64(match.Deaths)
				}
				seenMatches[m.ID] = match
			} else {
				// not seen we insert the match into the map
				m.Seen++
				seenMatches[m.ID] = m
			}
		}
	}

	// fmt.Printf("seen matches %+v", seenMatches)

	// loop through matches that we "collected"
	for id, match := range seenMatches {
		// if we have seen the match on every team
		if match.Seen == rules.TeamSize && !match.Checked {
			// add the games that we have seen to the total games
			t.Total.Games = append(t.Total.Games, match)
			// update the teams total stats
			t.Total.Kills += match.Kills
			t.Total.Deaths += match.Deaths
			t.Total.Headshots += match.Headshots
			if t.Total.Deaths == 0 {
				t.Total.KD = float64(t.Total.Kills)
			} else {
				t.Total.KD = (float64(t.Total.Kills) / float64(t.Total.Deaths))
			}
			t.Total.DamageDone += match.DamageDone
			t.Total.DamageTaken += match.DamageTaken
			if match.Placement == 1 {
				t.Total.Wins++
			}
			t.Total.PlacementPoints += match.Score
			t.Total.GamesPlayed++

		} else {
			// if we have not seen it the right amount of times, yeet it
			delete(seenMatches, id)
		}
	}
	// update the combined points total after we loop through all the matches
	t.Total.CombinedPoints = t.Total.PlacementPoints + t.Total.Kills

	// get the teams best matches
	t.Best.Games = sortMatches(t.Total.Games, rules.BestGamesNum)
	fmt.Printf("saved %d best matches for team %s\n", len(t.Best.Games), t.Name)

	// update the teams best stats
	// initialize an empty best struct
	best := Best{}
	fmt.Printf("")
	for _, match := range t.Best.Games {
		best.Kills += match.Kills
		best.Deaths += match.Deaths
		best.Headshots += match.Headshots
		best.DamageDone += match.DamageDone
		best.DamageTaken += match.DamageTaken
		if match.Placement == 1 {
			best.Wins++
		}
		best.PlacementPoints += match.Score
	}
	// update the KD without dividing by zero ever
	if best.Deaths == 0 {
		best.KD = float64(best.Kills)
	} else {
		best.KD = (float64(best.Kills) / float64(best.Deaths))
	}
	best.CombinedPoints = best.PlacementPoints + best.Kills
	// reassign the best.Games to the new best struct
	best.Games = t.Best.Games
	// reassign the best struct to the teams best
	t.Best = best
	fmt.Printf("updated all %d total matches stats for team %s with %d total kills\n", len(t.Total.Games), t.Name, t.Total.Kills)

	// go through the players and delete the matches that dont have the full team
	// // update the players total stats

	// get the players best matches
	// // update the players best stats
}

// sortMatches sorts the matches and returns the n best matches
func sortMatches(matches []Match, n int) []Match {
	sort.Sort(ByScore(matches))
	// slice the array if more than n
	if len(matches) > n {
		matches = matches[:n]
	}
	return matches
}
