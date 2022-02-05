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
			halfCell := float64(field.cellSize / 2)
			cellX, cellY := float64(x*field.cellSize)+halfCell, float64(y*field.cellSize)+halfCell

			if field.arrows {
				DrawArrow(c, cellX, cellY, vector.Angle(), halfCell)
			}
		}
	}

	return c
}

func DrawArrow(c *gg.Context, x float64, y float64, angle float64, length float64) {

	// Initialize context stack
	c.Push()

	// Arrow
	c.Translate(x, y)
	c.Rotate(-angle)
	c.DrawLine(0, 0, length, 0)
	c.Stroke()

	// Arrow Head
	arowHeadSize := 2.
	c.DrawLine(length, 0, length-arowHeadSize, -arowHeadSize)
	c.DrawLine(length-arowHeadSize, -arowHeadSize, length-arowHeadSize, arowHeadSize)
	c.DrawLine(length-arowHeadSize, arowHeadSize, length, 0)
	c.Stroke()

	// Dispose context stack
	c.Pop()
}
