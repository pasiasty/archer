package server

import (
	"fmt"
	"sync"
	"time"
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

// GetGame gets game with given ID.
func (gm *GameManager) GetGame(gameID string) (*Game, error) {
	game, ok := gm.games[gameID]
	if !ok {
		return nil, fmt.Errorf("failed to find game: %s", gameID)
	}
	return game, nil
}

// CreateGame creates new game and host user, returns host user.
func (gm *GameManager) CreateGame() *User {
	gm.mux.Lock()
	defer gm.mux.Unlock()

	newGameID := selectNewKey(unifyGamesMap(gm.games))
	newGame := CreateGame(newGameID)
	gm.games[newGameID] = newGame

	// removing games automatically after 3 days
	go func() {
		<-time.After(3 * 24 * time.Microsecond)
		gm.removeGame(newGameID)
	}()

	newUser := newGame.AddHostUser()

	return newUser
}

func (gm *GameManager) removeGame(gameID string) {
	delete(gm.games, gameID)
}

// JoinGame joins to game and creates new client user.
func (gm *GameManager) JoinGame(gameID string) (*User, error) {
	game, err := gm.GetGame(gameID)
	if err != nil {
		return nil, err
	}

	return game.AddClientUser(), nil
}

// AddPlayer adds another player to selected user and game.
func (gm *GameManager) AddPlayer(gameID, userID string) error {
	game, err := gm.GetGame(gameID)
	if err != nil {
		return err
	}

	return game.AddPlayer(userID)
}

// RemovePlayer removes player from selected user and game.
func (gm *GameManager) RemovePlayer(gameID, userID string) error {
	game, err := gm.GetGame(gameID)
	if err != nil {
		return err
	}

	return game.RemovePlayer(userID)
}

func unifyGamesMap(m map[string]*Game) map[string]bool {
	res := map[string]bool{}

	for k := range m {
		res[k] = true
	}

	return res
}
