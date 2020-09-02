package server

import (
	"math"
	"math/rand"
)

// Point contains coordinates.
type Point struct {
	X, Y float32
}

// RandomPoint returns random point.
func RandomPoint(edgeDistance float32) *Point {
	x := rand.Intn(maxX-2*int(edgeDistance)) + int(edgeDistance)
	y := rand.Intn(maxY-2*int(edgeDistance)) + int(edgeDistance)

	return &Point{
		X: float32(x),
		Y: float32(y),
	}
}

// Distance returns distance between two points.
func (p *Point) Distance(other *Point) float32 {
	diffX := p.X - other.X
	diffY := p.Y - other.Y
	return float32(math.Pow(float64(diffX*diffX+diffY*diffY), 0.5))
}

// Length returns length.
func (p *Point) Length() float32 {
	return p.Distance(&Point{X: 0, Y: 0})
}

// Normalize returns normalized point.
func (p *Point) Normalize() *Point {
	l := p.Length()
	return &Point{X: p.X / l, Y: p.Y / l}
}

// CopyWithSameAlpha returns copy of the point with same alpha, but different length.
func (p *Point) CopyWithSameAlpha(length float32) *Point {
	n := p.Normalize()
	return &Point{
		X: n.X * length,
		Y: n.Y * length,
	}
}
