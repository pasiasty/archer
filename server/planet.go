package server

// Planet contains planet information.
type Planet struct {
	Location *Point
	Radius   float32
	Mass     float32
}

// CreatePlanet creates new planet.
func CreatePlanet(location *Point, radius float32) *Planet {
	return &Planet{
		Location: location,
		Radius:   radius,
		Mass:     radius * radius * radius,
	}
}
