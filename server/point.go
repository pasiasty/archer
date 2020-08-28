package server

import (
	"math"
	"math/rand"
)

// Point contains coordinates.
type Point struct {
	x, y float32
}

// RandomPoint returns random point.
func RandomPoint() *Point {
	x := rand.Intn(maxX)
	y := rand.Intn(maxY)

	return &Point{
		x: float32(x),
		y: float32(y),
	}
}

// Distance returns distance between two points.
func (p *Point) Distance(other *Point) float32 {
	diffX := p.x - other.x
	diffY := p.y - other.y
	return float32(math.Pow(float64(diffX*diffX+diffY*diffY), 0.5))
}
