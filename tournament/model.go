package tournament

import (
	"time"
	"zleague/api/models"
)

// Tournament struct holds the information needed to start a tournament.
// TeamMates is an array of Activision Usernames.
type Tournament struct {
	ID        string
	StartTime time.Time
	EndTime   time.Time
	Teams     []models.Team
}

// TeamBasic holds a simple struct of what a team consists of.
type TeamBasic struct {
	Teamname  string
	Teammates []string
	Start     time.Time
	End       time.Time
	Division  string
}
