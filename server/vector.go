package server

import (
	"math"
	"math/rand"
)

// Vector contains coordinates.
type Vector struct {
	X, Y float32
}

// RandomVector returns random vector.
func RandomVector(edgeDistance float32) *Vector {
	x := rand.Intn(maxX-2*int(edgeDistance)) + int(edgeDistance)
	y := rand.Intn(maxY-2*int(edgeDistance)) + int(edgeDistance)

	return &Vector{
		X: float32(x),
		Y: float32(y),
	}
}

// Distance returns distance between two vectors.
func (p *Vector) Distance(other *Vector) float32 {
	diffX := p.X - other.X
	diffY := p.Y - other.Y
	return float32(math.Pow(float64(diffX*diffX+diffY*diffY), 0.5))
}

// Length returns length.
func (p *Vector) Length() float32 {
	return p.Distance(&Vector{X: 0, Y: 0})
}

// Normalize returns normalized point.
func (p *Vector) Normalize() *Vector {
	l := p.Length()
	return &Vector{X: p.X / l, Y: p.Y / l}
}

// CopyWithSameAlpha returns copy of the vector with same alpha, but different length.
func (p *Vector) CopyWithSameAlpha(length float32) *Vector {
	n := p.Normalize()
	return &Vector{
		X: n.X * length,
		Y: n.Y * length,
	}
}
