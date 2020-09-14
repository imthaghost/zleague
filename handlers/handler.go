package handlers

import (
	"zleague/api/tournament"

	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	db      *mongo.Database
	manager *tournament.TournamentManager
}

func NewHandler(db *mongo.Database, manager *tournament.TournamentManager) *Handler {
	return &Handler{
		db,
		manager,
	}
}
