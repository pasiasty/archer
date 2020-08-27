package server

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddPlayer(t *testing.T) {
	u := CreateUser("game", "user", "a", false)

	if diff := cmp.Diff([]string{"a"}, u.Players); diff != "" {
		t.Errorf("Wrong value of players, diff: %s", diff)
	}
	u.AddPlayer("b")

	if diff := cmp.Diff([]string{"a", "b"}, u.Players); diff != "" {
		t.Errorf("Wrong value of players, diff: %s", diff)
	}

	u.AddPlayer("c")

	if diff := cmp.Diff([]string{"a", "b", "c"}, u.Players); diff != "" {
		t.Errorf("Wrong value of players, diff: %s", diff)
	}
}

func TestRemovePlayer(t *testing.T) {
	u := CreateUser("game", "user", "a", false)
	u.AddPlayer("b")
	u.AddPlayer("c")

	if diff := cmp.Diff([]string{"a", "b", "c"}, u.Players); diff != "" {
		t.Errorf("Wrong value of players, diff: %s", diff)
	}
	if removed, _ := u.RemovePlayer(); removed != "c" {
		t.Errorf("Wrong name of removed player, want: %s got: %s", "c", removed)
	}

	if diff := cmp.Diff([]string{"a", "b"}, u.Players); diff != "" {
		t.Errorf("Wrong value of players, diff: %s", diff)
	}

	if removed, _ := u.RemovePlayer(); removed != "b" {
		t.Errorf("Wrong name of removed player, want: %s got: %s", "b", removed)
	}

	if diff := cmp.Diff([]string{"a"}, u.Players); diff != "" {
		t.Errorf("Wrong value of players, diff: %s", diff)
	}

	if _, err := u.RemovePlayer(); err == nil {
		t.Errorf("Removing should return error")
	}
}
