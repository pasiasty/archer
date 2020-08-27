package server

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	adjectives = []string{
		"Moon",
		"Chubby",
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
		"Pretty",
	}

	nouns = []string{
		"Moon",
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
		"Kiwi",
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

// GetUser gets selected user.
func (g *Game) GetUser(userID string) (*User, error) {
	user, ok := g.users[userID]
	if !ok {
		return nil, fmt.Errorf("failed to get user %s from game %s", userID, g.gameID)
	}
	return user, nil
}

// GetUsersList gets the list of game users.
func (g *Game) GetUsersList() []*PublicUser {
	res := []*PublicUser{}

	for _, u := range g.users {
		res = append(res, u.ConstructPublicUser())
	}

	return res
}

// AddPlayer ads another player to selected user.
func (g *Game) AddPlayer(userID string) error {
	user, err := g.GetUser(userID)
	if err != nil {
		return err
	}

	if user.Ready() {
		return fmt.Errorf("can't add player to ready user: %s", user.UserID)
	}

	g.mux.Lock()
	defer g.mux.Unlock()

	username := g.generateUsername()
	g.usernamesToIDs[username] = user.UserID
	user.AddPlayer(username)
	return nil
}

// RemovePlayer removes player from selected user.
func (g *Game) RemovePlayer(userID string) error {
	user, err := g.GetUser(userID)
	if err != nil {
		return err
	}

	if user.Ready() {
		return fmt.Errorf("can't remove player from ready user: %s", user.UserID)
	}

	g.mux.Lock()
	defer g.mux.Unlock()

	name, err := user.RemovePlayer()
	if err != nil {
		return err
	}
	delete(g.usernamesToIDs, name)
	return nil
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

// MarkUserReady marks selected user as ready.
func (g *Game) MarkUserReady(userID string) error {
	user, err := g.GetUser(userID)
	if err != nil {
		return err
	}
	user.MarkReady()
	return nil
}
