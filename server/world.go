package server

import (
	"errors"
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
	extraPlanets := 2 + rand.Intn(2)
	numOfPlanets := len(players) + extraPlanets
	res := &World{}

	fullnesRatio := FullnessRatio(len(players))

	for done := 0; done < numOfPlanets; {
		if p, err := res.addPlanet(fullnesRatio); err != nil {
			done = 0
			res.planets = nil
			continue
		} else {
			res.planets = append(res.planets, p)
			done++
		}
	}

	rand.Shuffle(len(players), func(i, j int) { players[i], players[j] = players[j], players[i] })
	for idx, p := range players {
		res.players = append(res.players, CreatePlayer(p, idx, res.planets[idx], rand.Float32()*2*math.Pi))
	}

	return res
}

func (w *World) addPlanet(fullnesRatio float32) (*Planet, error) {
	for counter := 0; counter < 32; counter++ {
		newRadius := minRadius + rand.Float32()*(maxRadius-minRadius)
		newRadius *= fullnesRatio
		newPoint := RandomPoint(newRadius + minPlanetDistance*fullnesRatio/2)

		wrong := false

		for _, p := range w.planets {
			dist := newPoint.Distance(p.Location)
			if dist < (newRadius + p.Radius + minPlanetDistance*fullnesRatio) {
				wrong = true
				break
			}
		}

		if !wrong {
			return CreatePlanet(newPoint, newRadius), nil
		}
	}
	return nil, errors.New("tried too many times")
}

// GetPublicWorld constructs PublicWorld from World.
func (w *World) GetPublicWorld() *PublicWorld {
	return &PublicWorld{
		Planets: w.planets,
	}
}
