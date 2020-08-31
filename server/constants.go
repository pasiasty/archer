package server

import (
	"math"
)

const maxPlayers = 12

const maxX = 1920.0
const maxY = 1080.0
const minRadius = 40.0
const maxRadius = 80.0
const minPlanetDistance = 130.0

// FullnessRatio magnifies distances and sizes according to number of players (the lower the bigger).
func FullnessRatio(players int) float32 {
	return float32(math.Pow(maxPlayers/float64(players), 0.4))
}

func init() {
	maxEdgeDistance := FullnessRatio(2) * (maxRadius + minPlanetDistance)
	if maxX < (2*maxEdgeDistance) || maxY < (2*maxEdgeDistance) {
		panic("no room for spreading the planets.")
	}
}
