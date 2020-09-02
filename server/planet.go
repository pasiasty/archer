package server

import "math/rand"

// Planet contains planet information.
type Planet struct {
	Location   Vector
	Radius     float32
	Mass       float32
	ResourceID int
	PlanetID   int
}

// CreatePlanet creates new planet.
func CreatePlanet(id int, location Vector, radius float32) *Planet {
	return &Planet{
		PlanetID:   id,
		Location:   location,
		Radius:     radius,
		Mass:       radius * radius * radius,
		ResourceID: rand.Intn(256),
	}
}
