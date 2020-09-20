package db

import (
	"context"
	"fmt"
	"log"
	"zleague/api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect will connect to the dev/production database
func Connect() *mongo.Database {
	dbConfig := config.GetDBConfig()

	auth := options.Credential{
		Username: dbConfig.Username,
		Password: dbConfig.Password,
	}
	uri := options.Client().ApplyURI(dbConfig.URI).SetAuth(auth)

	fmt.Println(uri)
	client, err := mongo.Connect(context.TODO(), uri)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("zleague")
}

// ConnectTest will connect to the test database
func ConnectTest() *mongo.Database {
	dbConfig := config.GetTestDBConfig()

	auth := options.Credential{
		Username: dbConfig.Username,
		Password: dbConfig.Password,
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbConfig.URI).SetAuth(auth))
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("zleague-test")
}
