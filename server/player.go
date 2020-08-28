package server

// Player is a game player.
type Player struct {
	name   string
	planet *Planet
	alpha  float32
}

// CreatePlayer creates new player.
func CreatePlayer(name string, planet *Planet, alpha float32) *Player {
	return &Player{
		name:   name,
		planet: planet,
		alpha:  alpha,
	}
}
