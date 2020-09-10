package tournament

import (
	"time"
	"zleague/backend_v2/models"

	"github.com/robfig/cron"
)

// Tournament struct holds the information needed to start a tournament.
// TeamMates is an array of Activision Usernames.
type Tournament struct {
	StartTime time.Time
	EndTime   time.Time
	Teams     []models.Team
	Cron      *cron.Cron
}

// TeamBasic holds a simple struct of what a team consists of.
type TeamBasic struct {
	Teamname  string
	Teammates []string
	Start     time.Time
	End       time.Time
	Division  string
}
