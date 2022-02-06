package core

import (
	"math"

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

func (field *VectorField) GetCell(x int, y int) *vectors.Vector2 {
	ix := math.Min(float64(field.columns-1), math.Max(0, float64(x)))
	iy := math.Min(float64(field.rows-1), math.Max(0, float64(y)))
	return field.vectors[int(iy)][int(ix)]
}

func (field *VectorField) SetCell(x int, y int, vector *vectors.Vector2) {
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
			adjustedX := cartesianX / float64(columns) * scale
			adjustedY := cartesianY / float64(rows) * scale
			vec := formula(vectors.NewVector2(adjustedX, adjustedY)).Normalize()

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
