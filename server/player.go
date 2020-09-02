package server

import "math"

// Player is a game player.
type Player struct {
	name     string
	planet   *Planet
	alpha    float32
	colorIdx int
}

// PublicPlayer is a Player structure passed to frontend.
type PublicPlayer struct {
	Name     string
	PlanetID int
	Alpha    float32
	ColorIdx int
}

// CreatePlayer creates new player.
func CreatePlayer(name string, colorIdx int, planet *Planet, alpha float32) *Player {
	return &Player{
		name:     name,
		planet:   planet,
		alpha:    alpha,
		colorIdx: colorIdx,
	}
}

// GetPublicPlayer returns Player structure to be passed to frontend.
func (p *Player) GetPublicPlayer() *PublicPlayer {
	return &PublicPlayer{
		Name:     p.name,
		PlanetID: p.planet.PlanetID,
		Alpha:    p.alpha,
		ColorIdx: p.colorIdx,
	}
}

// Coordinates returns global coordinates of the player.
func (p *Player) Coordinates() Vector {
	res := *p.planet.Location
	res.X += float32(float64(p.planet.Radius+playerShootHeight) * math.Sin(float64(p.alpha)))
	res.Y -= float32(float64(p.planet.Radius+playerShootHeight) * math.Cos(float64(p.alpha)))

	return res
}
