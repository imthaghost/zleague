package tournament

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

// NewTournament returns a Tournament instance
// teams should be a dictionary, where the key value is the team name, and the value is a string array of Activision ID's
func NewTournament(t map[string]TeamBasic, startTime time.Time, endTime time.Time) Tournament {
	teams := createTeams(t)
	return Tournament{
		StartTime: startTime,
		EndTime:   endTime,
		Teams:     teams,
	}
}

// StartTournamentWithDB starts the tournament from the Tournament struct instance and stores into a local mongodb
func (tournament *Tournament) StartTournamentWithDB() *cron.Cron {
	// default to every 10 minutes
	schedule := "@every 10m"
	// create new cron instance
	c := cron.New()
	// // db context
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// // db client
	// client := database.Connect(ctx)
	// // db handler
	// db := database.GetDatabase(client)

	fmt.Println("Starting:")
	// call function every 10 minutes
	c.AddFunc(schedule, func() {
		fmt.Println("Current time " + time.Now().Format(time.RFC3339))
		// we stop the cron job 30 minutes after the tournament endtime
		if time.Now().After(tournament.EndTime.Add(time.Minute * time.Duration(30))) {
			// graceful kill
			fmt.Println("stopping cron")
			c.Stop()
		}
		//iterate over team map

	})
	// start the jobs
	c.Start()
	// bind the cron to the tournament struct so we can stop it
	tournament.Cron = c
	// return it as well in case we need it
	return c
}
