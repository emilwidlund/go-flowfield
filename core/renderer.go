package core

import (
	"math"

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
	c.SetHexColor("#0000ff")

	for y, row := range field.vectors {
		for x, vector := range row {
			halfCell := field.cellSize / 2
			cellX, cellY := x*field.cellSize+halfCell, y*field.cellSize+halfCell
			DrawArrow(c, cellX, cellY, vector.Angle(), halfCell)
		}
	}
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

func DrawCurves(c *gg.Context, field *VectorField) {
	d := float64(field.width) / math.Sqrt((float64(field.curveCount)*float64(field.height))/float64(field.width))
	x0 := 0.                    // bbox min
	y0 := 0.                    // bbox min
	x1 := float64(field.width)  // bbox max
	y1 := float64(field.height) // bbox max
	r := d * 0.8                // min distance between points
	k := 15                     // max attempts to add neighboring point

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
