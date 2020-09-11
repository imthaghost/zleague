package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"zleague/backend_v2/tournament"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// try parsing start time
	start, err := time.Parse(time.RFC3339, "2020-09-08T09:20:00+00:00")
	if err != nil {
		log.Fatal(err)
	}
	// try parsing end time
	end, err := time.Parse(time.RFC3339, "2020-09-08T14:00:00+00:00")
	if err != nil {
		log.Fatal(err)
	}
	teams := tournament.Create(start, end)

	t := tournament.NewTournament(teams, start, end)

	t.UpdateTeam()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Database("tournaments").Collection("tournament").InsertOne(context.TODO(), t)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(res)
}
