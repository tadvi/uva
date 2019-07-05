package filter

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}

	a = Ints(a, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(a)
}
