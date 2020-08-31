package server

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/gobuffalo/buffalo"
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
	started        bool
	users          map[string]*User
	usernamesToIDs map[string]string
	world          *World
	maxPlayers     int
}

type gameOpts struct {
	maxPlayers int
}

// GameOption might be used for configuring game.
type GameOption func(*gameOpts)

// MaxPlayers overrides default max number of game players.
func MaxPlayers(max int) GameOption {
	return func(o *gameOpts) {
		o.maxPlayers = max
	}
}

// CreateGame creates new instance of the game.
func CreateGame(gameID string, opts ...GameOption) *Game {
	opt := gameOpts{
		maxPlayers: maxPlayers,
	}

	for _, o := range opts {
		o(&opt)
	}

	return &Game{
		gameID:         gameID,
		users:          map[string]*User{},
		usernamesToIDs: map[string]string{},
		maxPlayers:     opt.maxPlayers,
	}
}

// Started tells whether game has started.
func (g *Game) Started() bool {
	return g.started
}

// Start starts the game.
func (g *Game) Start(c buffalo.Context, userID string) error {
	if g.host.UserID != userID {
		return c.Error(http.StatusForbidden, fmt.Errorf("user: %s is not a host and can't start the game: %s", userID, g.gameID))
	}
	g.mux.Lock()
	defer g.mux.Unlock()

	if g.started {
		return nil
	}

	var players []string
	for p := range g.usernamesToIDs {
		players = append(players, p)
	}

	g.world = CreateWorld(players)
	g.started = true
	return nil
}

// GetWorld is used for getting the world state.
func (g *Game) GetWorld(c buffalo.Context) (*PublicWorld, error) {
	if !g.Started() {
		return nil, c.Error(http.StatusForbidden, fmt.Errorf("cannot get world for not started game: %s", g.gameID))
	}
	return g.world.GetPublicWorld(), nil
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
func (g *Game) GetUser(c buffalo.Context, userID string) (*User, error) {
	user, ok := g.users[userID]
	if !ok {
		return nil, c.Error(http.StatusNotFound, fmt.Errorf("failed to get user %s from game %s", userID, g.gameID))
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
func (g *Game) AddPlayer(c buffalo.Context, userID string) error {
	user, err := g.GetUser(c, userID)
	if err != nil {
		return err
	}

	if user.Ready() {
		return c.Error(http.StatusForbidden, fmt.Errorf("can't add player to ready user: %s", user.UserID))
	}

	if len(g.usernamesToIDs) >= g.maxPlayers {
		return c.Error(http.StatusForbidden, fmt.Errorf("can't add more then %d players", g.maxPlayers))
	}

	g.mux.Lock()
	defer g.mux.Unlock()

	username := g.generateUsername()
	g.usernamesToIDs[username] = user.UserID
	user.AddPlayer(username)
	return nil
}

// RemovePlayer removes player from selected user.
func (g *Game) RemovePlayer(c buffalo.Context, userID string) error {
	user, err := g.GetUser(c, userID)
	if err != nil {
		return err
	}

	if user.Ready() {
		return c.Error(http.StatusForbidden, fmt.Errorf("can't remove player from ready user: %s", user.UserID))
	}

	g.mux.Lock()
	defer g.mux.Unlock()

	name, err := user.RemovePlayer(c)
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
func (g *Game) MarkUserReady(c buffalo.Context, userID string) error {
	user, err := g.GetUser(c, userID)
	if err != nil {
		return err
	}
	user.MarkReady()
	return nil
}
