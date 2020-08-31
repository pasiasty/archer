package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gobuffalo/buffalo"
)

// User contains all information of single user.
type User struct {
	GameID   string
	UserID   string
	Username string
	Players  []string
	IsHost   bool
	ready    bool
	mux      sync.Mutex
}

// PublicUser contains public information of the user to be shared with other users.
type PublicUser struct {
	Username string
	Players  []string
	Ready    bool
	IsHost   bool
}

// CreateUser creates new user.
func CreateUser(gameID, userID, username string, isHost bool) *User {
	return &User{
		GameID:   gameID,
		UserID:   userID,
		Username: username,
		Players:  []string{username},
		IsHost:   isHost,
	}
}

// AddPlayer adds another player to the user.
func (u *User) AddPlayer(name string) {
	u.mux.Lock()
	defer u.mux.Unlock()
	u.Players = append(u.Players, name)
}

// RemovePlayer adds another player to the user.
func (u *User) RemovePlayer(c buffalo.Context) (string, error) {
	u.mux.Lock()
	defer u.mux.Unlock()

	if len(u.Players) <= 1 {
		return "", c.Error(http.StatusForbidden, fmt.Errorf("no extra players on user: %s", u.UserID))
	}
	lastPlayer := u.Players[len(u.Players)-1]
	u.Players[len(u.Players)-1] = ""
	u.Players = u.Players[:len(u.Players)-1]
	return lastPlayer, nil
}

// HasPlayer tells whether user contains specific player.
func (u *User) HasPlayer(player string) bool {
	for _, p := range u.Players {
		if p == player {
			return true
		}
	}

	return false
}

// StoreToCookie stores all user relevant information into cookie.
func (u *User) StoreToCookie(c buffalo.Context) {
	SetCookie(c, "game_id", u.GameID)
	SetCookie(c, "user_id", u.UserID)
	SetCookie(c, "username", u.Username)
	SetCookie(c, "is_host", fmt.Sprintf("%v", u.IsHost))
}

// ConstructPublicUser returns public information of the user.
func (u *User) ConstructPublicUser() *PublicUser {
	return &PublicUser{
		Username: u.Username,
		Players:  u.Players,
		Ready:    u.ready,
		IsHost:   u.IsHost,
	}
}

// MarkReady sets ready to true.
func (u *User) MarkReady() {
	u.ready = true
}

// Ready tells whether user is ready.
func (u *User) Ready() bool {
	return u.ready
}
