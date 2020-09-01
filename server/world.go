package server

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"sync"

	"github.com/gobuffalo/buffalo"
)

// World contains information of all players in the world.
type World struct {
	players              []*Player
	planets              []*Planet
	currentPlayerIdx     int
	returnedTrajectories int
	mux                  sync.Mutex
	currentTrajectory    *Trajectory
	numUsers             int
}

// PublicWorld is structure that will be returned to frontend.
type PublicWorld struct {
	Planets       []*Planet
	Players       []*PublicPlayer
	CurrentPlayer *PublicPlayer
}

// CreateWorld creates new world.
func CreateWorld(numUsers int, players []string) *World {
	extraPlanets := 2 + rand.Intn(2)
	numOfPlanets := len(players) + extraPlanets
	res := &World{
		numUsers: numUsers,
	}

	fullnesRatio := FullnessRatio(len(players))

	for done := 0; done < numOfPlanets; {
		if p, err := res.generatePlanet(done, fullnesRatio); err != nil {
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

func (w *World) generatePlanet(newPlanetID int, fullnesRatio float32) (*Planet, error) {
	for counter := 0; counter < 128; counter++ {
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
			return CreatePlanet(newPlanetID, newPoint, newRadius), nil
		}
	}
	return nil, errors.New("tried too many times")
}

// GetPublicWorld constructs PublicWorld from World.
func (w *World) GetPublicWorld() *PublicWorld {
	players := []*PublicPlayer{}

	for _, p := range w.players {
		players = append(players, p.GetPublicPlayer())
	}

	return &PublicWorld{
		Planets:       w.planets,
		Players:       players,
		CurrentPlayer: players[w.currentPlayerIdx],
	}
}

// GetPollTurn returns active player state.
func (w *World) GetPollTurn() PollTurn {
	return PollTurn{
		CurrentPlayer:      w.players[w.currentPlayerIdx].name,
		CurrentPlayerAlpha: w.players[w.currentPlayerIdx].alpha,
		ShotPerformed:      !(w.returnedTrajectories == 0),
	}
}

// MovePlayer sets new alpha for current player.
func (w *World) MovePlayer(c buffalo.Context, player string, newAlpha float32) error {
	w.mux.Lock()
	defer w.mux.Unlock()
	currentPlayer := w.players[w.currentPlayerIdx]
	if currentPlayer.name != player {
		return c.Error(http.StatusForbidden, fmt.Errorf("player: %s is not an active one", player))
	}
	currentPlayer.alpha = newAlpha
	return nil
}

// Shoot performs shot for selected player.
func (w *World) Shoot(c buffalo.Context, player string, shot Point) (*Trajectory, error) {
	w.mux.Lock()
	defer w.mux.Unlock()
	currentPlayer := w.players[w.currentPlayerIdx]
	if currentPlayer.name != player {
		return nil, c.Error(http.StatusForbidden, fmt.Errorf("player: %s is not an active one", player))
	}
	t := w.generateTrajectory(shot)
	w.returnedTrajectories = 1
	w.currentTrajectory = t

	return t, nil
}

func (w *World) generateTrajectory(shot Point) *Trajectory {
	return &Trajectory{}
}

// GetTrajectory returns current trajectory.
func (w *World) GetTrajectory() *Trajectory {
	w.mux.Lock()
	defer w.mux.Unlock()
	w.returnedTrajectories++
	t := w.currentTrajectory
	if w.returnedTrajectories == w.numUsers {
		w.returnedTrajectories = 0
		w.currentTrajectory = nil
	}
	return t
}
