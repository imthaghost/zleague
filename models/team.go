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
		// convert matches to match struct and store on the player, then update the matches
		// in the seenMatches histogram, in order to store on the team properly
		t.Players[i].getMatches(matches, &rules).updateMatches(&seenMatches)
	}
	// if the lenght of the seenMatches map is 0, then no new matches have been played, return
	if len(seenMatches) == 0 {
		return
	}
	// Update the total stats struct using the histogram while conforming to the rules
	t.updateStats(&seenMatches, rules)
}

// Method to update the total and best stats on a team
func (t *Team) updateStats(seenMatches *map[string]Match, rules Rules) {
	// update the total stats on a team and then update the best stats.
	t.updateTotal(seenMatches, rules).updateBest()
	// go through the players and delete the matches that dont have the full team
	// // update the players total stats
	t.updatePlayersStats(seenMatches, rules)
}

// updateTotal will update the total points on a team
func (t *Team) updateTotal(seenMatches *map[string]Match, rules Rules) *Team {
	// loop through matches that we "collected"
	for _, match := range *seenMatches {
		// if we have seen the match on every team
		if match.Seen == rules.TeamSize {
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

		}
	}
	// update the combined points total after we loop through all the matches
	t.Total.CombinedPoints = t.Total.PlacementPoints + t.Total.Kills

	// get the teams best matches
	t.Best.Games = sortMatches(t.Total.Games, rules.BestGamesNum)
	return t
}

// updateBest will update the best struct on a team
func (t *Team) updateBest() {
	// update the teams best stats
	// initialize an empty best struct
	best := Best{}
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
}

// Updates the players best stats, based off of the matches stored on their struct
func (t *Team) updatePlayersStats(seenMatches *map[string]Match, rules Rules) {
	// for every player, call the updateBest method
	// can add concurrency for efficiency here
	for i := range t.Players {
		t.Players[i].updateStats(seenMatches, rules)
	}
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
