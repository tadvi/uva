package matrix

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {

	m := MakeInts(4, 4, func(i, j int) int {
		return i + j
	})

	Apply(m, func(i, j int) {
		fmt.Println(m[i])
	})

}
