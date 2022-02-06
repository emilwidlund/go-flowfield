package core

import (
	"math"

	"github.com/emilwidlund/esmerelda/vectors"
	"github.com/fogleman/gg"
	"github.com/fogleman/poissondisc"
)

func Draw(field *VectorField) *gg.Context {
	c := gg.NewContext(field.width, field.height)
	c.SetHexColor("#ffffff")
	c.Clear()
	c.SetHexColor("#000000")

	for y, row := range field.vectors {
		for x, vector := range row {
			halfCell := field.cellSize / 2
			cellX, cellY := x*field.cellSize+halfCell, y*field.cellSize+halfCell

			if field.arrows {
				DrawArrow(c, cellX, cellY, vector.Angle(), halfCell)
			}
		}
	}

	if !field.arrows {
		DrawSimulation(c, field)
	}

	return c
}

func DrawArrow(c *gg.Context, x int, y int, angle float64, length int) {
	c.Push()

	// Arrow
	c.Translate(float64(x), float64(y))
	c.Rotate(-angle)
	c.DrawLine(0, 0, float64(length), 0)
	c.Stroke()

	// Arrow Head
	arowHeadSize := 4.
	ahx := float64(length) - arowHeadSize
	c.DrawLine(float64(length), 0, ahx, -arowHeadSize)
	c.DrawLine(ahx, arowHeadSize, float64(length), 0)
	c.Stroke()

	c.Pop()
}

func DrawSimulation(c *gg.Context, field *VectorField) {
	const SIMULATION_COUNT = 500

	d := float64(field.width) / math.Sqrt((float64(SIMULATION_COUNT)*float64(field.height))/float64(field.width))
	x0 := 0.                    // bbox min
	y0 := 0.                    // bbox min
	x1 := float64(field.width)  // bbox max
	y1 := float64(field.height) // bbox max
	r := d * 0.8                // min distance between points
	k := 15                     // max attempts to add neighboring point

	points := poissondisc.Sample(x0, y0, x1, y1, r, k, nil)

	for _, p := range points {
		DrawCurve(c, field, int(p.X), int(p.Y))
	}

}

func DrawCurve(c *gg.Context, field *VectorField, x int, y int) {
	const STEP_LENGTH = 1
	const NUM_STEPS = 500

	c.Push()

	p := vectors.NewVector2(float64(x), float64(y))
	q := vectors.NewVector2(float64(x), float64(y))
	n := NUM_STEPS >> 1

	curve := make([]vectors.Vector2, 0)

	for n > 0 {
		n--
		angle := field.GetAngle(p.X, p.Y)
		v := vectors.NewVector2(1, 0).Rotate(-angle).Scale(STEP_LENGTH)
		p.Add(v)
		curve = append(curve, *p)
	}

	ReverseCurve(curve)
	n = NUM_STEPS - (NUM_STEPS >> 1)
	for n > 0 {
		n--
		angle := field.GetAngle(q.X, q.Y)
		v := vectors.NewVector2(-1, 0).Rotate(-angle).Scale(STEP_LENGTH)
		q.Add(v)
		curve = append(curve, *q)
	}

	for _, v := range curve {
		c.LineTo(v.X, v.Y)
	}

	c.Stroke()

	c.Pop()

}

func Sgn(a float64) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}

func ShortAngleDist(a0 float64, a1 float64) float64 {
	max := math.Pi * 2
	da := float64(Sgn(a1-a0)) * math.Mod(math.Abs(a1-a0), max)

	return float64(Sgn(a1-a0))*math.Mod((2*math.Abs(da)), max) - da
}

func AngleLerp(a0 float64, a1 float64, t float64) float64 {
	return a0 + ShortAngleDist(a0, a1)*t
}

func ReverseCurve(curve []vectors.Vector2) []vectors.Vector2 {
	for i, j := 0, len(curve)-1; i < j; i, j = i+1, j-1 {
		curve[i], curve[j] = curve[j], curve[i]
	}

	return curve
}
