package server

import (
	"fmt"
	"sync"
)

var (
	gm = &GameManager{
		games: map[string]*Game{},
	}
)

// GameManager manages basic operations with games.
type GameManager struct {
	games map[string]*Game
	mux   sync.Mutex
}

// CreateGameManager creates new GameManager.
func CreateGameManager() *GameManager {
	return &GameManager{
		games: map[string]*Game{},
	}
}

// CreateGame creates new game and host user, returns host user.
func (gm *GameManager) CreateGame() *User {
	gm.mux.Lock()
	defer gm.mux.Unlock()

	newGameID := selectNewKey(unifyGamesMap(gm.games))
	newGame := CreateGame(newGameID)
	gm.games[newGameID] = newGame

	newUser := newGame.AddHostUser()

	return newUser
}

// JoinGame joins to game and creates new client user.
func (gm *GameManager) JoinGame(gameID string) (*User, error) {
	game, ok := gm.games[gameID]
	if !ok {
		return nil, fmt.Errorf("failed to find game: %s", gameID)
	}

	return game.AddClientUser(), nil
}

func unifyGamesMap(m map[string]*Game) map[string]bool {
	res := map[string]bool{}

	for k := range m {
		res[k] = true
	}

	return res
}
