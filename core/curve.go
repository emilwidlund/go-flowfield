package core

import (
	"github.com/emilwidlund/esmerelda/vectors"
)

type Curve struct {
	path []vectors.Vector2
}

func NewCurve(field *VectorField, x float64, y float64) *Curve {
	curve := &Curve{
		path: make([]vectors.Vector2, 0),
	}

	p := vectors.NewVector2(x, y)
	q := vectors.NewVector2(x, y)
	n := field.numSteps >> 1

	for n > 0 {
		n--
		angle := field.GetAngle(p.X, p.Y)
		v := vectors.NewVector2(1, 0).Rotate(-angle).Scale(field.stepSize)
		p.Add(v)
		curve.AddSegment(p)
	}

	curve.Reverse()
	n = field.numSteps - (field.numSteps >> 1)
	for n > 0 {
		n--
		angle := field.GetAngle(q.X, q.Y)
		v := vectors.NewVector2(-1, 0).Rotate(-angle).Scale(field.stepSize)
		q.Add(v)
		curve.AddSegment(q)
	}

	return curve
}

func (curve *Curve) AddSegment(segment *vectors.Vector2) *Curve {
	curve.path = append(curve.path, *segment)

	return curve
}

func (curve *Curve) Reverse() *Curve {
	for i, j := 0, len(curve.path)-1; i < j; i, j = i+1, j-1 {
		curve.path[i], curve.path[j] = curve.path[j], curve.path[i]
	}

	return curve
}
