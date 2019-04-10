package main

import (
	"fmt"
	"io"
	"os"
)

type cache map[int]int

func (memo cache) count(n int) int {
	cnt := 1
	for n != 1 {
		if v, ok := memo[n]; ok {
			return v
		}

		if n%2 != 0 {
			n = 3*n + 1
		} else {
			n = n / 2
		}
		cnt++
	}

	memo[n] = cnt
	return cnt
}

func run(w io.Writer) {
	r, _ := os.Open("input.txt")
	defer r.Close()

	var memo = cache{}

	var a, b int
	for {
		if _, err := fmt.Fscan(r, &a, &b); err != nil {
			break
		}
		if a > b {
			a, b = b, a
		}

		for i := a; i <= b; i++ {
			memo.count(i)
		}

		fmt.Fprintf(w, "%d %d\n", a, b)
	}
}

func main() {
	run(os.Stdout)
}
