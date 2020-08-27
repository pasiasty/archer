package server

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	adjectives = []string{
		"Chubby",
		"Honky",
		"Brave",
		"Silent",
		"Sneaky",
		"Incredible",
		"Smart",
		"Clean",
		"Pure",
		"Vicious",
		"Envy",
		"Angry",
		"Complicated",
	}

	nouns = []string{
		"Panda",
		"Horse",
		"Elephant",
		"Man",
		"Duck",
		"Tiger",
		"Lion",
		"Crocodile",
		"Dog",
		"Cat",
		"Kivi",
		"Koala",
		"Bear",
	}
)

// Game contains all information relevant for game.
type Game struct {
	mux            sync.Mutex
	host           *User
	gameID         string
	users          map[string]*User
	usernamesToIDs map[string]string
}

// CreateGame creates new instance of the game.
func CreateGame(gameID string) *Game {
	return &Game{
		gameID:         gameID,
		users:          map[string]*User{},
		usernamesToIDs: map[string]string{},
	}
}

// AddHostUser adds new user as host.
func (g *Game) AddHostUser() *User {
	return g.addUser(true)
}

// AddClientUser adds new user as client.
func (g *Game) AddClientUser() *User {
	return g.addUser(false)
}

// HasUser tells whether game contains the user.
func (g *Game) HasUser(userID string) bool {
	_, ok := g.users[userID]
	return ok
}

func (g *Game) generateUsername() string {
	for {
		adjIdx, nounIdx := rand.Intn(len(adjectives)), rand.Intn(len(nouns))
		newUsername := fmt.Sprintf("%s %s", adjectives[adjIdx], nouns[nounIdx])
		if _, ok := g.usernamesToIDs[newUsername]; !ok {
			return newUsername
		}
	}
}

func (g *Game) addUser(asHost bool) *User {
	g.mux.Lock()
	defer g.mux.Unlock()

	newUser := CreateUser(
		g.gameID,
		selectNewKey(unifyUsersMap(g.users)),
		g.generateUsername(),
		asHost,
	)
	g.users[newUser.UserID] = newUser
	g.usernamesToIDs[newUser.Username] = newUser.UserID

	if asHost {
		g.host = newUser
	}

	return newUser
}

func unifyUsersMap(m map[string]*User) map[string]bool {
	res := map[string]bool{}

	for k := range m {
		res[k] = true
	}

	return res
}
