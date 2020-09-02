package server

import (
	"fmt"
	"math"
	"testing"
)

func Test_World_CreateWorld(t *testing.T) {
	for numPlayers := 2; numPlayers < maxPlayers; numPlayers += 4 {
		players := []string{}

		for i := 0; i < numPlayers; i++ {
			players = append(players, "")
		}
		for i := 0; i < 16; i++ {
			t.Run(fmt.Sprintf("%d players variant: %d", numPlayers, i), func(t *testing.T) {
				CreateWorld(0, players)
			})
		}
	}
}

func Test_World_vectorToAngle(t *testing.T) {
	for _, tc := range []struct {
		p     Point
		angle float32
	}{{
		p:     Point{X: 1, Y: 0},
		angle: 0,
	}, {
		p:     Point{X: 1, Y: 1},
		angle: math.Pi * 1 / 4,
	}, {
		p:     Point{X: 0, Y: 1},
		angle: math.Pi * 2 / 4,
	}, {
		p:     Point{X: -1, Y: 1},
		angle: math.Pi * 3 / 4,
	}, {
		p:     Point{X: -1, Y: 0},
		angle: math.Pi * 4 / 4,
	}, {
		p:     Point{X: -1, Y: -1},
		angle: math.Pi * 5 / 4,
	}, {
		p:     Point{X: 0, Y: -1},
		angle: math.Pi * 6 / 4,
	}, {
		p:     Point{X: 1, Y: -1},
		angle: math.Pi * 7 / 4,
	}} {
		t.Run(fmt.Sprintf("%v_%v", tc.p, tc.angle), func(t *testing.T) {
			if res := vectorToAngle(tc.p); !floatCompare(res, tc.angle, 0.001) {
				t.Errorf("wrong angle value, want: %v, got: %v", tc.angle, res)
			}
		})
	}
}
