package config

import "os"

type DBConfig struct {
	URI      string
	Username string
	Password string
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		URI:      os.Getenv("DB_URI"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

func GetTestDBConfig() *DBConfig {
	return &DBConfig{
		URI:      os.Getenv("DB_TEST_URI"),
		Username: os.Getenv("DB_TEST_USERNAME"),
		Password: os.Getenv("DB_TEST_PASSWORD"),
	}
}
