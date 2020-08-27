package actions

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"sync"
)

var (
	gm = &gameManager{
		games: map[string]*game{},
	}
)

type gameManager struct {
	games map[string]*game
	mux   sync.Mutex
}

func unifyGamesMap(m map[string]*game) map[string]bool {
	res := map[string]bool{}

	for k := range m {
		res[k] = true
	}

	return res
}

func unifyUsersMap(m map[string]*user) map[string]bool {
	res := map[string]bool{}

	for k := range m {
		res[k] = true
	}

	return res
}

func selectNewKey(m map[string]bool) string {
	h := sha1.New()
	for {
		io.WriteString(h, fmt.Sprintf("%d", rand.Intn(10000)))
		res := base64.URLEncoding.EncodeToString(h.Sum(nil))

		if _, ok := m[res]; !ok {
			return res
		}
	}
}

func (gm *gameManager) createGame() (string, string) {
	gm.mux.Lock()
	defer gm.mux.Unlock()
	newGameID := selectNewKey(unifyGamesMap(gm.games))
	users := map[string]*user{}
	newUserID := selectNewKey(unifyUsersMap(users))
	host := &user{
		userID: newUserID,
	}
	users[newUserID] = host

	gm.games[newGameID] = &game{
		host:  host,
		users: users,
	}

	return newGameID, newUserID
}

func (gm *gameManager) joinGame(gameID string) (string, error) {
	game, ok := gm.games[gameID]
	if !ok {
		return "", fmt.Errorf("failed to find game: %s", gameID)
	}

	game.mux.Lock()
	defer game.mux.Unlock()

	newUserID := selectNewKey(unifyUsersMap(game.users))
	newUser := &user{
		userID: newUserID,
	}
	game.users[newUserID] = newUser
	return newUserID, nil
}

type game struct {
	mux   sync.Mutex
	host  *user
	users map[string]*user
}

type user struct {
	userID string
}
