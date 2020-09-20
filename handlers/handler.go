package handlers

import (
	"zleague/api/tournament"

	"go.mongodb.org/mongo-driver/mongo"
)

// Handler is a struct that all the handler (routes) are a method of so we can dependency inject a database and tournament manager
type Handler struct {
	db      *mongo.Database
	manager *tournament.Manager
}

// New will return a new handler struct.
func New(db *mongo.Database, manager *tournament.Manager) *Handler {
	return &Handler{
		db,
		manager,
	}
}
