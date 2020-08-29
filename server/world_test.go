package server

import (
	"fmt"
	"testing"
)

func Test_World_CreateWorld(t *testing.T) {
	for numPlayers := 2; numPlayers < maxPlayers; numPlayers += 4 {
		players := []string{}

		for i := 0; i < numPlayers; i++ {
			players = append(players, "")
		}
		for i := 0; i < 16; i++ {
			t.Run(fmt.Sprintf("%d players, variant: %d", numPlayers, i), func(t *testing.T) {
				CreateWorld(players)
			})
		}
	}
}
