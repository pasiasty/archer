package actions

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"
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

func (gm *gameManager) createGame() string {
	h := sha1.New()
	io.WriteString(h, fmt.Sprintf("%d", rand.Intn(10000)))
	for {
		gm.mux.Lock()
		res := base64.URLEncoding.EncodeToString(h.Sum(nil))
		if _, ok := gm.games[res]; !ok {
			host := &user{userID: 0}
			gm.games[res] = &game{
				host:  host,
				users: []*user{host},
			}
			gm.mux.Unlock()
			return res
		}
		gm.mux.Unlock()
		time.Sleep(time.Microsecond)
	}
}

func (gm *gameManager) joinGame(gameID string) int {
	game, ok := gm.games[gameID]
	if !ok {
		return -1
	}

	game.mux.Lock()
	defer game.mux.Unlock()

	newUserID := len(game.users)
	game.users = append(game.users, &user{userID: newUserID})
	return newUserID
}

type game struct {
	mux   sync.Mutex
	host  *user
	users []*user
}

type user struct {
	userID int
}
