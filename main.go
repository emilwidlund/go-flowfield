package main

import (
	"math"

	"github.com/aquilax/go-perlin"
	"github.com/emilwidlund/esmerelda/core"
	"github.com/emilwidlund/esmerelda/vectors"
)

func noiseFormula() core.Formula {
	p := perlin.NewPerlin(2, 2, 3, 0)

	return func(vector *vectors.Vector2) *vectors.Vector2 {
		n := p.Noise2D(vector.X/40, vector.Y/40) * 10

		return vectors.NewVector2(math.Cos(n+10), math.Sin(n))
	}
}

/*
func circleFormula() core.Formula {
	return func(vector *vectors.Vector2) *vectors.Vector2 {
		return vectors.NewVector2(0, -1).Rotate((vector.Y / 20) * math.Pi)
	}
}

func sineFormula() core.Formula {
	return func(vector *vectors.Vector2) *vectors.Vector2 {
		return vectors.NewVector2(math.Sin(vector.Y), math.Sin(vector.X))
	}
}

func customFormula() core.Formula {
	return func(vector *vectors.Vector2) *vectors.Vector2 {
		return vectors.NewVector2((vector.X*vector.X - vector.Y*vector.Y), vector.X*vector.Y)
	}
}*/

func main() {
	field := core.NewVectorField(1920, 1080, noiseFormula(), 30, 5000, 2, 500, false)
	context := core.Draw(field)

	context.SavePNG("test.png")
}
