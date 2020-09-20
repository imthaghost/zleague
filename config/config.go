package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Did not load env variables from file. This is fine if you are in production/ci environment.")
	}
}
