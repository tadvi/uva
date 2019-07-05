package functool

import (
	"fmt"
	"testing"
)

func TestPermutate(t *testing.T) {

	Permutate([]int{1, 2, 3}, func(slice []int) {
		fmt.Println(slice)
	})
}
