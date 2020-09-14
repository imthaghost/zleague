package tournament

import "net/http"

// TournamentManager is designed to manage multiple tournaments at once.
// It does this through a database, where it stores current tournaments and information about them.
// You can create new tournaments, delete tournaments, and more.
type TournamentManager struct {
	tournaments map[string]Tournament // represents all current tournaments
	client      http.Client
}

// NewTournamentManager will create a new instance of tournament manager.
// By default, it will load tournaments that are currently in the database so that they can be interacted with.
func NewTournamentManager() *TournamentManager {
	// TODO: logic
	return &TournamentManager{}
}

// NewTourament is designed to create a new tournament, and then save it to the struct and the database and return it
func (t *TournamentManager) NewTourament() Tournament {
	// TODO: logic
	return Tournament{}
}

// GetTournament will get a tournament from memory or w/e
// GET /tournament/:id
// GET /tournament/:id/teams/:id
func (t *TournamentManager) GetTournament(id string) Tournament {
	// TODO: logic
	return Tournament{}
}

func (t *TournamentManager) AllTournaments() []Tournament {
	// TODO: logic
	return []Tournament{}
}
