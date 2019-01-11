// UVa 846 - Steps

package main

import (
	"fmt"
	"os"
)

func solve(n int) int {
	cnt, step := 0, 1
	for n > 0 {
		cnt++
		if n -= step; n == 0 {
			break
		}
		if cnt%2 == 0 {
			step++
		}
	}
	return cnt
}

func main() {
	in, _ := os.Open("846.in")
	defer in.Close()
	out, _ := os.Create("846.out")
	defer out.Close()

	var n, a, b int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d%d", &a, &b)
		fmt.Fprintln(out, solve(b-a))
	}
}
