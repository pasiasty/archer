package server

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
)

// User contains all information of single user.
type User struct {
	GameID   string
	UserID   string
	Username string
	IsHost   bool
	ready    bool
}

// PublicUser contains public information of the user to be shared with other users.
type PublicUser struct {
	Username string
	Ready    bool
	IsHost   bool
}

// CreateUser creates new user.
func CreateUser(gameID, userID, username string, isHost bool) *User {
	return &User{
		GameID:   gameID,
		UserID:   userID,
		Username: username,
		IsHost:   isHost,
	}
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
		Ready:    u.ready,
		IsHost:   u.IsHost,
	}
}
