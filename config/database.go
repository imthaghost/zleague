package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DBConfig represents the values needed for a database connection
type DBConfig struct {
	URI      string
	Username string
	Password string
}

// GetDBConfig will return the default database connection.
func GetDBConfig() *DBConfig {
	return &DBConfig{
		URI:      os.Getenv("DB_URI"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

// GetTestDBConfig will return the test database config.
func GetTestDBConfig() *DBConfig {
	// attempt to load env in test environment
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Did not load environment variables from .env. This is normal if you are not in development.")
	}

	return &DBConfig{
		URI:      os.Getenv("DB_TEST_URI"),
		Username: os.Getenv("DB_TEST_USERNAME"),
		Password: os.Getenv("DB_TEST_PASSWORD"),
	}
}
