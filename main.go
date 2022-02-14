package main

import (
	"math"

	"github.com/aquilax/go-perlin"
)

func noiseFormula() Formula {
	p := perlin.NewPerlin(2, 2, 3, 0)

	return func(vector *Vector) *Vector {
		n := p.Noise2D(vector.X/40, vector.Y/40) * 10

		return NewVector(math.Cos(n+10), math.Sin(n))
	}
}

/*
func identityFormula() core.Formula {
	return func(vector *Vector) *Vector {
		return vector
	}
}

func noiseFormula() core.Formula {
	p := perlin.NewPerlin(2, 2, 3, 0)

	return func(vector *Vector) *Vector {
		n := p.Noise2D(vector.X/40, vector.Y/40) * 10

		return NewVector(math.Cos(n+10), math.Sin(n))
	}
}

func circleFormula() core.Formula {
	return func(vector *Vector) *Vector {
		return NewVector(0, -1).Rotate((vector.Y / 20) * math.Pi)
	}
}

func sineFormula() core.Formula {
	return func(vector *Vector) *Vector {
		return NewVector(math.Sin(vector.Y), math.Sin(vector.X))
	}
}

func customFormula() core.Formula {
	return func(vector *Vector) *Vector {
		return NewVector((vector.X*vector.X - vector.Y*vector.Y), vector.X*vector.Y)
	}
}*/

func main() {
	field := NewVectorField(1920, 1080, noiseFormula(), 30, 500, 1, 1000, true)
	context := Draw(field)

	context.SavePNG("test.png")
}
