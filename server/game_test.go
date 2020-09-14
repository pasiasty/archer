package server

import (
	"context"
	"testing"

	"github.com/gobuffalo/buffalo"
)

func Test_Game_Start(t *testing.T) {
	c := &buffalo.DefaultContext{
		Context: context.Background(),
	}

	g := CreateGame("a")
	host := g.AddHostUser()
	client := g.AddClientUser()

	if host.UserID == client.UserID || host.Username == client.Username {
		t.Errorf("users should have different IDs and usernames. Host: %v, Client: %v", host, client)
	}
	if err := g.Start(c, client.UserID); err == nil {
		t.Errorf("game should not let to be started by client")
	}
	if err := g.Start(c, host.UserID); err != nil {
		t.Errorf("game should let to be started by host")
	}
	if !g.started {
		t.Errorf("Game should be started by now.")
	}
}

func Test_Game_GetUser(t *testing.T) {
	c := &buffalo.DefaultContext{
		Context: context.Background(),
	}

	g := CreateGame("a")
	host := g.AddHostUser()

	u, err := g.GetUser(c, host.UserID)
	if err != nil {
		t.Errorf("failed to get user.")
	}
	if u != host {
		t.Errorf("got wrong user. want: %+v got: %+v", host, u)
	}
	if _, err := g.GetUser(c, "abc"); err == nil {
		t.Error("getting non-existing user should fail.")
	}
}

func Test_Game_AddingRemovingPlayers(t *testing.T) {
	c := &buffalo.DefaultContext{
		Context: context.Background(),
	}

	g := CreateGame("a")
	host := g.AddHostUser()

	if err := g.RemovePlayer(c, "abc"); err == nil {
		t.Errorf("removing player for nonexisting user should fail")
	}

	if err := g.RemovePlayer(c, host.UserID); err == nil {
		t.Errorf("removing player for 1 player user should not be possible")
	}

	if err := g.AddPlayer(c, host.UserID); err != nil {
		t.Errorf("failed to add player: %v", err)
	}

	if err := g.AddPlayer(c, host.UserID); err != nil {
		t.Errorf("failed to add player: %v", err)
	}

	if err := g.AddPlayer(c, host.UserID); err != nil {
		t.Errorf("failed to add player: %v", err)
	}

	if numPlayers := len(g.GetUsersList()[0].Players); numPlayers != 4 {
		t.Errorf("wrong number of numPlayers, want: %d got: %d", 4, numPlayers)
	}

	if err := g.RemovePlayer(c, host.UserID); err != nil {
		t.Errorf("failed to remove player: %v", err)
	}

	if err := g.RemovePlayer(c, host.UserID); err != nil {
		t.Errorf("failed to remove player: %v", err)
	}

	if numPlayers := len(g.GetUsersList()[0].Players); numPlayers != 2 {
		t.Errorf("wrong number of numPlayers, want: %d got: %d", 2, numPlayers)
	}

	host.MarkReady()

	if err := g.AddPlayer(c, host.UserID); err == nil {
		t.Errorf("adding player for ready user should fail")
	}

	if err := g.RemovePlayer(c, host.UserID); err == nil {
		t.Errorf("removing player for ready user should fail")
	}
}

func Test_Game_ExceedingMaxPlayers(t *testing.T) {
	c := &buffalo.DefaultContext{
		Context: context.Background(),
	}

	g := CreateGame("a", WithMaxPlayers(2))
	host := g.AddHostUser()

	if err := g.AddPlayer(c, host.UserID); err != nil {
		t.Errorf("failed to add player: %v", err)
	}
	if err := g.AddPlayer(c, host.UserID); err == nil {
		t.Errorf("exceeding max players should fail.")
	}
}
