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
			cellX := float64(x*field.cellSize + (field.cellSize / 2))
			cellY := float64(y * field.cellSize)
			c.DrawLine(cellX, cellY, cellX, float64(int(cellY)+field.cellSize))
			c.Rotate(vector.Angle())
		}
	}

	c.Stroke()

	return c
}
