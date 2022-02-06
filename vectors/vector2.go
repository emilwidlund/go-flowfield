package vectors

import (
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

func NewVector2(x float64, y float64) *Vector2 {
	return &Vector2{x, y}
}

func (a *Vector2) Add(b *Vector2) *Vector2 {
	a.X += b.X
	a.Y += b.Y

	return a
}

func (vector *Vector2) AddScalar(scalar float64) *Vector2 {
	vector.X += scalar
	vector.Y += scalar

	return vector
}

func (a *Vector2) Subtract(b *Vector2) *Vector2 {
	a.X -= b.X
	a.Y -= b.Y

	return a
}

func (vector *Vector2) SubtractScalar(scalar float64) *Vector2 {
	vector.X -= scalar
	vector.Y -= scalar

	return vector
}

func (a *Vector2) Multiply(b *Vector2) *Vector2 {
	a.X *= b.X
	a.Y *= b.Y

	return a
}

func (vector *Vector2) MultiplyScalar(scalar float64) *Vector2 {
	vector.X *= scalar
	vector.Y *= scalar

	return vector
}

func (a *Vector2) Divide(b *Vector2) *Vector2 {
	a.X /= b.X
	a.Y /= b.Y

	return a
}

func (vector *Vector2) DivideScalar(scalar float64) *Vector2 {
	vector.X /= scalar
	vector.Y /= scalar

	return vector
}

func (vector *Vector2) Scale(scalar float64) *Vector2 {
	vector.X *= scalar
	vector.Y *= scalar

	return vector
}

func (vector *Vector2) Normalize() *Vector2 {
	magnitude := vector.Magnitude()

	if magnitude == 0 {
		return vector.DivideScalar(1)
	} else {
		return vector.DivideScalar(magnitude)
	}
}

func (vector *Vector2) Rotate(angle float64) *Vector2 {
	c, s := math.Cos(angle), math.Sin(angle)

	vector.X = vector.X*c - vector.Y*s
	vector.Y = vector.X*s + vector.Y*c

	return vector
}

func (vector *Vector2) Magnitude() float64 {
	return math.Sqrt(math.Pow(vector.X, 2) + math.Pow(vector.Y, 2))
}

func (a *Vector2) DistanceTo(b *Vector2) float64 {
	return a.Subtract(b).Magnitude()
}

func (vector *Vector2) Angle() float64 {
	return math.Atan2(vector.Y, vector.X)
}
