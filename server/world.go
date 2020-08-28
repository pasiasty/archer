package server

import (
	"math"
	"math/rand"
)

// World contains information of all players in the world.
type World struct {
	players          []*Player
	planets          []*Planet
	currentPlayerIdx int
}

// PublicWorld is structure that will be returned to frontend.
type PublicWorld struct {
	Planets []*Planet
}

// CreateWorld creates new world.
func CreateWorld(players []string) *World {
	extraPlanets := 3 + rand.Intn(5)
	numOfPlanets := len(players) + extraPlanets
	res := &World{}

	for idx := 0; idx < numOfPlanets; idx++ {
		res.addPlanet()
	}

	rand.Shuffle(len(players), func(i, j int) { players[i], players[j] = players[j], players[i] })
	for idx, p := range players {
		res.players = append(res.players, CreatePlayer(p, res.planets[idx], rand.Float32()*2*math.Pi))
	}

	return res
}

func (w *World) addPlanet() {
	for {
		newPoint := RandomPoint()
		newRadius := minRadius + rand.Float32()*(maxRadius-minRadius)

		for _, p := range w.planets {
			dist := newPoint.Distance(p.Location)
			if dist < (newRadius + p.Radius + minPlanetDistance) {
				break
			}
		}
		w.planets = append(w.planets, CreatePlanet(newPoint, newRadius))
		return
	}
}

// GetPublicWorld constructs PublicWorld from World.
func (w *World) GetPublicWorld() *PublicWorld {
	return &PublicWorld{
		Planets: w.planets,
	}
}
