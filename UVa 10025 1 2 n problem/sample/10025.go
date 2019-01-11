// UVa 10025 - The ? 1 ? 2 ? ... ? n = k problem

package main

import (
	"fmt"
	"os"
)

func solve(k int) int {
	if k < 0 {
		k = -k
	}
	i, sum := 0, 0
	for sum < k || (sum-k)%2 != 0 {
		i++
		sum += i
	}
	return i
}

func main() {
	in, _ := os.Open("10025.in")
	defer in.Close()
	out, _ := os.Create("10025.out")
	defer out.Close()

	var kase, k int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d", &k)
		fmt.Fprintln(out, solve(k))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
