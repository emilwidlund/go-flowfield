package main

import (
	"math"

	"github.com/emilwidlund/esmerelda/core"
	"github.com/emilwidlund/esmerelda/vectors"
)

func main() {
	formula := func(vector *vectors.Vector2) *vectors.Vector2 {
		x := math.Sin(vector.X) + math.Sin(vector.Y)
		y := math.Sin(vector.X) - math.Sin(vector.Y)

		return vectors.NewVector2(x, y)
	}

	field := core.NewVectorField(800, 600, formula)
	context := core.Draw(field)

	context.SavePNG("test.png")
}
