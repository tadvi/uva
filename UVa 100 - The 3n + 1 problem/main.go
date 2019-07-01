package main

import (
	"fmt"
	"os"
)

func count(memo map[int]int, n int) int {
	if v, ok := memo[n]; ok {
		return v
	}

	var v int
	if n%2 == 0 {
		v = count(memo, n/2) + 1
	} else {
		v = count(memo, 3*n+1) + 1
	}

	memo[n] = v
	return v
}

func run() {
	r, _ := os.Open("input.txt")
	defer r.Close()

	memo := map[int]int{1: 1}

	var a, b int
	for {
		if _, err := fmt.Fscan(r, &a, &b); err != nil {
			break
		}
		if a > b {
			a, b = b, a
		}

		max := 0
		for i := a; i <= b; i++ {
			if v := count(memo, i); max < v {
				max = v
			}
		}

		fmt.Println(a, b, max)
	}
}

func main() {
	run()
}
