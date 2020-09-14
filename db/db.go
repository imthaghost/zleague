package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	auth := options.Credential{
		Username: "root",
		Password: "AVeryStrongPassword1234",
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(auth))
	if err != nil {
		log.Fatal(err)
	}

	return client
}
