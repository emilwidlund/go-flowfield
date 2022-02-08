package main

import (
	"math"
)

type Formula func(vector *Vector) *Vector

type VectorField struct {
	width      int
	height     int
	columns    int
	rows       int
	cellSize   int
	curveCount int
	stepSize   float64
	numSteps   int
	arrows     bool
	vectors    [][]*Vector
	formula    Formula
}

func NewVectorField(width int, height int, formula Formula, cellSize int, curveCount int, stepSize float64, numSteps int, arrows bool) *VectorField {
	columns := width / cellSize
	rows := height / cellSize

	v := make([][]*Vector, rows)

	for y := 0; y < rows; y++ {
		rowVectors := make([]*Vector, columns)

		for x := 0; x < columns; x++ {
			cartesianX, cartesianY := float64(x-columns/2), float64(rows/2-y)
			scale := 10.
			adjustedX := cartesianX / float64(columns) * scale
			adjustedY := cartesianY / float64(rows) * scale
			vec := formula(NewVector(adjustedX, adjustedY)).Normalize()

			rowVectors[x] = vec
		}

		v[y] = rowVectors
	}

	return &VectorField{
		width:      width,
		height:     height,
		columns:    columns,
		rows:       rows,
		cellSize:   cellSize,
		curveCount: curveCount,
		stepSize:   stepSize,
		numSteps:   numSteps,
		arrows:     arrows,
		vectors:    v,
		formula:    formula,
	}
}

func (field *VectorField) GetCell(x int, y int) *Vector {
	ix := math.Min(float64(field.columns-1), math.Max(0, float64(x)))
	iy := math.Min(float64(field.rows-1), math.Max(0, float64(y)))
	return field.vectors[int(iy)][int(ix)]
}

func (field *VectorField) SetCell(x int, y int, vector *Vector) {
	if x < field.columns && x >= 0 && y < field.rows && y >= 0 {
		field.vectors[y][x] = vector
	}
}

func (field *VectorField) GetCellIndex(x float64, y float64) (int, int) {
	return int(math.Floor(x / float64(field.cellSize))), int(math.Floor(y / float64(field.cellSize)))
}

func (field *VectorField) GetAngle(x float64, y float64) float64 {
	ix, iy := field.GetCellIndex(x, y)

	alphaX := math.Mod(x, float64(field.cellSize)) / float64(field.cellSize)
	alphaY := math.Mod(y, float64(field.cellSize)) / float64(field.cellSize)

	return AngleLerp(
		AngleLerp(field.GetCell(ix, iy).Angle(), field.GetCell(ix+1, iy).Angle(), alphaX),
		AngleLerp(field.GetCell(ix, iy+1).Angle(), field.GetCell(ix+1, iy+1).Angle(), alphaX),
		alphaY,
	)

}

func ShortAngleDist(a0 float64, a1 float64) float64 {
	max := math.Pi * 2
	da := float64(Sgn(a1-a0)) * math.Mod(math.Abs(a1-a0), max)

	return float64(Sgn(a1-a0))*math.Mod((2*math.Abs(da)), max) - da
}

func AngleLerp(a0 float64, a1 float64, t float64) float64 {
	return a0 + ShortAngleDist(a0, a1)*t
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
