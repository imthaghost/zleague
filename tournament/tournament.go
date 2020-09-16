package tournament

import (
	"time"
)

// NewTournament returns a Tournament instance
// teams should be a dictionary, where the key value is the team name, and the value is a string array of Activision ID's
func NewTournament(t map[string]TeamBasic, id string, startTime, endTime time.Time) Tournament {
	teams := createTeams(t)
	return Tournament{
		ID:        id,
		StartTime: startTime,
		EndTime:   endTime,
		Teams:     teams,
	}
}


