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
			// _, cellMidY := cellX+float64(field.cellSize/2), cellY+float64(field.cellSize/2)
			// _, cellEndY := float64(x*field.cellSize+field.cellSize), float64(y*field.cellSize+field.cellSize)

			c.Push()
			c.Translate(cellX, cellY)
			c.Rotate(-vector.Angle())
			c.DrawLine(0, 0, 10, 0)
			c.Stroke()
			c.DrawCircle(10, 0, 2)
			c.Fill()
			c.Pop()
		}
	}

	return c
}
