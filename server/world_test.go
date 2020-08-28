package server

import (
	"fmt"
	"testing"
)

func Test_World_CreateWorld(t *testing.T) {
	for numPlayers := 2; numPlayers < maxPlayers; numPlayers++ {
		fmt.Printf("generating for %d players\n", numPlayers)
		players := []string{}
		for i := 0; i < numPlayers; i++ {
			players = append(players, "")
		}
		for i := 0; i < 16; i++ {
			fmt.Printf("trial: %d\n", i)
			CreateWorld(players)
		}
	}
}
