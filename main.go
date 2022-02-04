package main

import (
	"github.com/emilwidlund/esmerelda/core"
	"github.com/emilwidlund/esmerelda/vectors"
)

func main() {
	formula := func(vector *vectors.Vector2) *vectors.Vector2 {
		x := .1 * vector.Y
		y := -0.2 * vector.Y

		return vectors.NewVector2(x, y)
	}

	field := core.NewVectorField(800, 600, formula)
	context := core.Draw(field)

	context.SavePNG("test.png")
}
