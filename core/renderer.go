package core

import (
	"math"

	"github.com/emilwidlund/esmerelda/vectors"
	"github.com/fogleman/gg"
	"github.com/fogleman/poissondisc"
)

func Draw(field *VectorField) *gg.Context {
	c := gg.NewContext(field.width, field.height)

	c.SetHexColor("#000033")
	c.Clear()

	if field.arrows {
		DrawArrows(c, field)
	} else {
		DrawCurves(c, field)
	}

	return c
}

func DrawArrows(c *gg.Context, field *VectorField) {
	for y, row := range field.vectors {
		for x, vector := range row {
			halfCell := field.cellSize / 2
			cellX, cellY := x*field.cellSize+halfCell, y*field.cellSize+halfCell
			DrawArrow(c, cellX, cellY, vector.Angle(), halfCell)
		}
	}
}

func DrawArrow(c *gg.Context, x int, y int, angle float64, length int) {
	r, g, b := GetColorByAngle(angle)
	c.SetRGB(r, g, b)

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

func DrawCurves(c *gg.Context, field *VectorField) {
	cc, w, h := float64(field.curveCount), float64(field.width), float64(field.height)

	d := w / math.Sqrt((cc*h)/w)
	x0 := 0.     // bbox min
	y0 := 0.     // bbox min
	x1 := w      // bbox max
	y1 := h      // bbox max
	r := d * 0.8 // min distance between points
	k := 15      // max attempts to add neighboring point

	points := poissondisc.Sample(x0, y0, x1, y1, r, k, nil)

	for i, p := range points {
		if i%4 == 1 {
			c.SetHexColor("#000044")
		} else {
			c.SetHexColor("#0000ff")
		}

		DrawCurve(c, field, p.X, p.Y)
	}
}

func DrawCurve(c *gg.Context, field *VectorField, x float64, y float64) {
	c.Push()

	for _, v := range NewCurve(field, x, y).path {
		r, g, b := GetColorByAngle(v.DistanceTo(vectors.NewVector2(x, y)))
		c.SetRGB(r, g, b)
		c.LineTo(v.X, v.Y)
	}

	c.Stroke()
	c.Pop()
}

func GetColorByAngle(angle float64) (r float64, g float64, b float64) {
	ra, ga, ba := math.Cos(angle), math.Cos(angle+20), math.Cos(angle-20)

	return Clamp(Normalize(ra)+.2, 0, 1), Clamp(Normalize(ga)+.2, 0, 1), Clamp(Normalize(ba)+.2, 0, 1)
}

func Clamp(value float64, min float64, max float64) float64 {
	return math.Min(math.Max(value, min), max)
}

func Normalize(value float64) float64 {
	return (value + 1) * .5
}
