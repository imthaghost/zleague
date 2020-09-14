package handlers

import (
	"context"
	"log"
	"net/http"
	"zleague/api/db"
	"zleague/api/tournament"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *Handler) GetTeams(c echo.Context) error {
	tournamentID := c.Param("id")

	var tournament tournament.Tournament
	err := db.Connect().Collection("tournaments").FindOne(context.TODO(), bson.M{"id": tournamentID}).Decode(&tournament)
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, tournament)
}
