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
func RandomVector(edgeDistance float32) Vector {
	x := rand.Intn(maxX-2*int(edgeDistance)) + int(edgeDistance)
	y := rand.Intn(maxY-2*int(edgeDistance)) + int(edgeDistance)

	return Vector{
		X: float32(x),
		Y: float32(y),
	}
}

// Distance returns distance between two vectors.
func (v Vector) Distance(other Vector) float32 {
	diffX := v.X - other.X
	diffY := v.Y - other.Y
	return float32(math.Pow(float64(diffX*diffX+diffY*diffY), 0.5))
}

// Length returns length.
func (v Vector) Length() float32 {
	return v.Distance(Vector{X: 0, Y: 0})
}

// Normalize returns normalized point.
func (v Vector) Normalize() Vector {
	l := v.Length()
	return Vector{X: v.X / l, Y: v.Y / l}
}

// CopyWithSameAlpha returns copy of the vector with same alpha, but different length.
func (v Vector) CopyWithSameAlpha(length float32) Vector {
	n := v.Normalize()
	return n.Mult(length)
}

// Add adds another vector to this one.
func (v Vector) Add(other Vector) Vector {
	res := Vector{X: v.X, Y: v.Y}
	res.X += other.X
	res.Y += other.Y
	return res
}

// Sub subtracts another vector to this one.
func (v Vector) Sub(other Vector) Vector {
	res := Vector{X: v.X, Y: v.Y}
	res.X -= other.X
	res.Y -= other.Y
	return res
}

// Mult multiplies vector by scalar.
func (v Vector) Mult(a float32) Vector {
	res := Vector{X: v.X, Y: v.Y}
	res.X *= a
	res.Y *= a
	return res
}
