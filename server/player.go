package server

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
