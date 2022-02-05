package core

import (
	"github.com/emilwidlund/esmerelda/vectors"
)

type Formula func(vector *vectors.Vector2) *vectors.Vector2

type VectorField struct {
	width    int
	height   int
	columns  int
	rows     int
	cellSize int
	arrows   bool
	vectors  [][]*vectors.Vector2
	formula  Formula
}

func NewVectorField(width int, height int, arrows bool, formula Formula) *VectorField {
	const CELL_SIZE = 30

	columns := width / CELL_SIZE
	rows := height / CELL_SIZE

	v := make([][]*vectors.Vector2, rows)

	for y := 0; y < rows; y++ {
		rowVectors := make([]*vectors.Vector2, columns)

		for x := 0; x < columns; x++ {
			cartesianX, cartesianY := float64(x-columns/2), float64(rows/2-y)
			scale := 10.
			vec := formula(vectors.NewVector2(cartesianX/float64(columns)*scale, cartesianY/float64(rows)*scale)).Normalize()

			rowVectors[x] = vec
		}

		v[y] = rowVectors
	}

	return &VectorField{
		width:    width,
		height:   height,
		columns:  columns,
		rows:     rows,
		cellSize: CELL_SIZE,
		arrows:   arrows,
		vectors:  v,
		formula:  formula,
	}
}
