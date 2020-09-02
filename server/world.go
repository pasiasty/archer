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
		newPoint := RandomVector(newRadius + minPlanetDistance*fullnesRatio/2)

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
func (w *World) Shoot(c buffalo.Context, player string, shot Vector) (*Trajectory, error) {
	w.mux.Lock()
	defer w.mux.Unlock()
	currentPlayer := w.players[w.currentPlayerIdx]
	if currentPlayer.name != player {
		return nil, c.Error(http.StatusForbidden, fmt.Errorf("player: %s is not an active one", player))
	}

	t := w.generateTrajectory(currentPlayer.name, currentPlayer.Coordinates(), shot)
	w.returnedTrajectories = 1
	w.currentTrajectory = t
	w.endTurnIfNeeded()

	return t, nil
}

func flipRadianIfNegative(alpha float32) float32 {
	if alpha < 0 {
		return alpha + 2*math.Pi
	}
	return alpha
}

func vectorToAngle(v Vector) float32 {
	if v.X >= 0 {
		return flipRadianIfNegative(float32(math.Atan(float64(v.Y / v.X))))
	}
	if v.Y >= 0 {
		return flipRadianIfNegative(math.Pi - float32(math.Atan(float64(-v.Y/v.X))))
	}
	return flipRadianIfNegative(math.Pi - float32(math.Atan(float64(-v.Y/v.X))))
}

func (w *World) applyGravity(pos, vel Vector) Vector {
	for _, p := range w.planets {
		dist := pos.Distance(p.Location)
		accLen := gravityConst * p.Mass / (dist * dist)
		accDir := p.Location.Sub(pos)
		acc := accDir.CopyWithSameAlpha(accLen)
		vel = vel.Add(acc.Mult(simulationTimeStep))
	}
	return vel
}

func (w *World) outsideBoundingBox(pos Vector) bool {
	if pos.X > 2*maxX || pos.X < -maxX || pos.Y > 4*maxY || pos.Y < -2*maxY {
		return true
	}
	return false
}

func (w *World) collidedWithPlanet(pos Vector) (bool, Vector) {
	for _, p := range w.planets {
		if pos.Distance(p.Location) <= p.Radius {
			pos = pos.Sub(p.Location).CopyWithSameAlpha(p.Radius).Add(p.Location)
			return true, pos
		}
	}
	return false, Vector{}
}

func (w *World) generateTrajectory(shooter string, start, shot Vector) *Trajectory {
	t := &Trajectory{}
	pos := start
	vel := shot.Mult(velScaling)

	arrowOffset := vel.CopyWithSameAlpha(arrowHalfLength)
	pos = pos.Add(arrowOffset)

	alpha := vectorToAngle(shot)

	for i := 0; i < maxSimulationSamples; i++ {
		t.ArrowStates = append(t.ArrowStates, ArrowState{
			Orientation: alpha,
			Position:    pos,
		})

		for _, p := range w.players {
			if p.name != shooter && p.Collision(pos) {
				t.CollidedWith = p.name
				return t
			}
		}

		pos = pos.Add(vel.Mult(simulationTimeStep))
		vel = w.applyGravity(pos, vel)
		alpha = vectorToAngle(vel)

		if w.outsideBoundingBox(pos) {
			return t
		}

		if collided, corrPos := w.collidedWithPlanet(pos); collided {
			t.ArrowStates = append(t.ArrowStates, ArrowState{
				Orientation: alpha,
				Position:    corrPos,
			})
			t.CollidedWith = "planet"
			return t
		}
	}
	return t
}

func (w *World) removeKilledPlayer(name string) {
	idxToRemove := -1
	for idx, player := range w.players {
		if player.name == name {
			idxToRemove = idx
			break
		}
	}

	if idxToRemove == -1 {
		return
	}

	if idxToRemove <= w.currentPlayerIdx {
		w.currentPlayerIdx--
	}

	copy(w.players[idxToRemove:], w.players[idxToRemove+1:]) // Shift a[i+1:] left one index.
	w.players[len(w.players)-1] = nil                        // Erase last element (write zero value).
	w.players = w.players[:len(w.players)-1]                 // Truncate slice.
}

// GetTrajectory returns current trajectory.
func (w *World) GetTrajectory() *Trajectory {
	w.mux.Lock()
	defer w.mux.Unlock()
	w.returnedTrajectories++
	t := w.currentTrajectory
	w.endTurnIfNeeded()
	return t
}

func (w *World) endTurnIfNeeded() {
	t := w.currentTrajectory
	if w.returnedTrajectories == w.numUsers {
		w.returnedTrajectories = 0
		w.currentTrajectory = nil
		if t.CollidedWith != "" && t.CollidedWith != "planet" {
			w.removeKilledPlayer(t.CollidedWith)
		}
		w.currentPlayerIdx = (w.currentPlayerIdx + 1) % len(w.players)
	}
}
