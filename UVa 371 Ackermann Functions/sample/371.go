// UVa 371 - Ackermann Functions

package main

import (
	"fmt"
	"os"
)

func calculate(i int) int {
	cache := make(map[int]int)
	if _, ok := cache[i]; !ok {
		if i%2 == 0 {
			if i/2 == 1 { // so that f(1)=3 as requested in problem
				return 1
			}
			cache[i] = 1 + calculate(i/2)
		} else {
			cache[i] = 1 + calculate(3*i+1)
		}
	}
	return cache[i]
}

func solve(l, h int) (int, int) {
	num, max := 0, 0
	for i := l; i <= h; i++ {
		n := calculate(i)
		if n > max {
			num, max = i, n
		}
	}
	return num, max
}

func main() {
	in, _ := os.Open("371.in")
	defer in.Close()
	out, _ := os.Create("371.out")
	defer out.Close()

	var l, h int
	for {
		if fmt.Fscanf(in, "%d%d", &l, &h); l == 0 && h == 0 {
			break
		}
		num, max := solve(l, h)
		fmt.Fprintf(out, "Between %d and %d, %d generates the longest sequence of %d values.\n", l, h, num, max)
	}
}
