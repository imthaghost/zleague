package main

import (
	"fmt"
	"log"
	"time"
	"zleague/api/db"
	"zleague/api/tournament"
)

func main() {
	now := time.Now()
	// try parsing start time
	start, err := time.Parse(time.RFC3339, "2020-09-11T01:50:00+00:00")
	if err != nil {
		log.Fatal(err)
	}
	// try parsing end time
	end, err := time.Parse(time.RFC3339, "2020-09-11T4:50:00+00:00")
	if err != nil {
		log.Fatal(err)
	}
	teams := tournament.Create(start, end)

	t := tournament.NewTournament(teams, start, end)
	client := db.Connect()
	t.Insert(client.Database("tournaments"))

	t.Update()

	t.UpdateInDB(client.Database("tournaments"))
	// res, err := client.Database("tournaments").Collection("tournament").InsertOne(context.TODO(), t)
	// if err != nil {
	// 	log.Println(err)
	// }

	fmt.Println(time.Now().Sub(now))

	// start the server
	// server := server.NewServer(nil)
	// server.Start(":8080")
}
