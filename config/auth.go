package config

import "os"

// AuthConfig will return the config for basic auth
type AuthConfig struct {
	Username string
	Password string
}

// GetAuthConfig will return the auth config
func GetAuthConfig() *AuthConfig {
	return &AuthConfig{
		Username: os.Getenv("SERVER_USERNAME"),
		Password: os.Getenv("SERVER_PASSWORD"),
	}
}
