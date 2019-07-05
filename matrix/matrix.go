package matrix

import "reflect"

func Apply(mat interface{}, fn func(row, col int)) {
	rv := reflect.ValueOf(mat)
	hl := rv.Len()
	for i := 0; i < hl; i++ {

		vl := rv.Index(i).Len()
		for j := 0; j < vl; j++ {
			fn(i, j)
		}
	}
}

func MakeInts(rows, cols int, fn func(row, col int) int) [][]int {
	mat := make([][]int, rows)
	for i := range mat {

		mat[i] = make([]int, cols)
		for j := range mat[i] {
			mat[i][j] = fn(i, j)
		}
	}
	return mat
}

func MakeFloat64s(rows, cols int, fn func(row, col int) float64) [][]float64 {
	mat := make([][]float64, rows)
	for i := range mat {

		mat[i] = make([]float64, cols)
		for j := range mat[i] {
			mat[i][j] = fn(i, j)
		}
	}
	return mat
}

func MakeStrings(rows, cols int, fn func(row, col int) string) [][]string {
	mat := make([][]string, rows)
	for i := range mat {

		mat[i] = make([]string, cols)
		for j := range mat[i] {
			mat[i][j] = fn(i, j)
		}
	}
	return mat
}
