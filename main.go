package main

import (
	"math"

	"github.com/aquilax/go-perlin"
	"github.com/emilwidlund/esmerelda/core"
	"github.com/emilwidlund/esmerelda/vectors"
)

func main() {

	p := perlin.NewPerlin(2, 2, 3, 0)

	formula := func(vector *vectors.Vector2) *vectors.Vector2 {
		// x := 1.
		// y := vector.Y*vector.Y - vector.Y

		n := p.Noise2D(vector.X/20, vector.Y/20) * 10

		return vectors.NewVector2(math.Cos(n), math.Sin(n))
	}

	field := core.NewVectorField(800, 600, formula)
	context := core.Draw(field)

	context.SavePNG("test.png")
}
