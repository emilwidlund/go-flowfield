package core

import (
	"github.com/fogleman/gg"
)

func Draw(field *VectorField) *gg.Context {
	c := gg.NewContext(field.width, field.height)
	c.SetHexColor("#ffffff")
	c.Clear()
	c.SetHexColor("#000000")

	for y, row := range field.vectors {
		for x, vector := range row {
			cellX, cellY := float64(x*field.cellSize), float64(y*field.cellSize)

			c.Push()
			c.Translate(cellX, cellY)
			c.Rotate(-vector.Angle())
			c.DrawLine(0, 0, 10, 0)
			c.Stroke()
			c.DrawLine(10, 0, 9, -1)
			c.DrawLine(9, -1, 9, 1)
			c.DrawLine(9, 1, 10, 0)
			c.Stroke()
			c.Pop()
		}
	}

	return c
}
