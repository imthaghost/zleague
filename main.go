package main

import (
	"zleague/api/server"
)

func main() {
	// now := time.Now()
	// // try parsing start time
	// start, err := time.Parse(time.RFC3339, "2020-09-11T01:50:00+00:00")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// // try parsing end time
	// end, err := time.Parse(time.RFC3339, "2020-09-11T4:50:00+00:00")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// teams := tournament.Create(start, end)
	//
	// t := tournament.NewTournament(teams, start, end)
	//
	// t.Update()
	//
	// client := db.Connect()
	//
	// res, err := client.Database("tournaments").Collection("tournament").InsertOne(context.TODO(), t)
	// if err != nil {
	//     log.Println(err)
	// }
	//
	// fmt.Println(res, time.Now().Sub(now))

	// start the server
	server := server.NewServer(nil)
	server.Start(":8080")
}
