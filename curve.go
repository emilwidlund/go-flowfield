package main

type Curve struct {
	path []Vector
}

func NewCurve(field *VectorField, x float64, y float64) *Curve {
	curve := &Curve{
		path: make([]Vector, 0),
	}

	p := NewVector(x, y)
	q := NewVector(x, y)
	n := field.numSteps >> 1

	for n > 0 {
		n--
		angle := field.GetAngle(p.X, p.Y)
		v := NewVector(1, 0).Rotate(-angle).Scale(field.stepSize)
		p.Add(v)
		curve.AddSegment(p)
	}

	curve.Reverse()
	n = field.numSteps - (field.numSteps >> 1)
	for n > 0 {
		n--
		angle := field.GetAngle(q.X, q.Y)
		v := NewVector(-1, 0).Rotate(-angle).Scale(field.stepSize)
		q.Add(v)
		curve.AddSegment(q)
	}

	return curve
}

func (curve *Curve) AddSegment(segment *Vector) *Curve {
	curve.path = append(curve.path, *segment)

	return curve
}

func (curve *Curve) Reverse() *Curve {
	for i, j := 0, len(curve.path)-1; i < j; i, j = i+1, j-1 {
		curve.path[i], curve.path[j] = curve.path[j], curve.path[i]
	}

	return curve
}
