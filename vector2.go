package main

import (
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func NewVector(x float64, y float64) *Vector {
	return &Vector{x, y}
}

func (a *Vector) Add(b *Vector) *Vector {
	a.X += b.X
	a.Y += b.Y

	return a
}

func (vector *Vector) AddScalar(scalar float64) *Vector {
	vector.X += scalar
	vector.Y += scalar

	return vector
}

func (a *Vector) Subtract(b *Vector) *Vector {
	a.X -= b.X
	a.Y -= b.Y

	return a
}

func (vector *Vector) SubtractScalar(scalar float64) *Vector {
	vector.X -= scalar
	vector.Y -= scalar

	return vector
}

func (a *Vector) Multiply(b *Vector) *Vector {
	a.X *= b.X
	a.Y *= b.Y

	return a
}

func (vector *Vector) MultiplyScalar(scalar float64) *Vector {
	vector.X *= scalar
	vector.Y *= scalar

	return vector
}

func (a *Vector) Divide(b *Vector) *Vector {
	a.X /= b.X
	a.Y /= b.Y

	return a
}

func (vector *Vector) DivideScalar(scalar float64) *Vector {
	vector.X /= scalar
	vector.Y /= scalar

	return vector
}

func (vector *Vector) Scale(scalar float64) *Vector {
	vector.X *= scalar
	vector.Y *= scalar

	return vector
}

func (vector *Vector) Normalize() *Vector {
	magnitude := vector.Magnitude()

	if magnitude == 0 {
		return vector.DivideScalar(1)
	} else {
		return vector.DivideScalar(magnitude)
	}
}

func (vector *Vector) Rotate(angle float64) *Vector {
	c, s := math.Cos(angle), math.Sin(angle)
	x, y := vector.X*c-vector.Y*s, vector.X*s+vector.Y*c

	vector.X = x
	vector.Y = y

	return vector
}

func (vector *Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(vector.X, 2) + math.Pow(vector.Y, 2))
}

func (a *Vector) DistanceTo(b *Vector) float64 {
	s := *a
	return s.Subtract(b).Magnitude()
}

func (vector *Vector) Angle() float64 {
	return math.Atan2(vector.Y, vector.X)
}
