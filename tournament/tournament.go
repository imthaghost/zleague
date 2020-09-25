package tournament

import (
	"log"
	"zleague/api/models"
)

// NewTournament returns a Tournament instance
// teams should be a dictionary, where the key value is the team name, and the value is a string array of Activision ID's
func NewTournament(id string, rules models.Rules, t map[string]models.TeamBasic) models.Tournament {
	// create the teams
	teams := createTeams(t)

	log.Println("Creating New Tournament... ID:", id)

	return models.Tournament{
		ID:    id,
		Teams: teams,
		Rules: rules,
	}
}
