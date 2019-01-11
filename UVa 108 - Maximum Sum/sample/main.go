package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMax(s [][]int) int {
	msf := math.MinInt32
	length := len(s)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {

			meh := 0
			for k := 0; k < length; k++ {
				tmp := s[j][k]
				if i > 0 {
					tmp -= s[i-1][k]
				}
				meh = max(meh+tmp, tmp)
				msf = max(msf, meh)
			}
		}

	}
	return msf
}

func prefixSum(n int, ra [][]int) [][]int {
	s := make([][]int, n)
	for i := range s {
		s[i] = make([]int, n)
		for j := range s[i] {
			s[i][j] = ra[i][j]
			if i > 0 {
				s[i][j] += s[i-1][j]
			}
		}
	}
	return s
}

func run(w io.Writer) {
	r, _ := os.Open("input.txt")
	defer r.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(r, "%d\n", &n); err != nil {
			break
		}

		ra := make([][]int, n)
		for i := range ra {
			ra[i] = make([]int, n)
			for j := range ra[i] {
				fmt.Fscan(r, &ra[i][j])
			}
		}
		fmt.Println(findMax(prefixSum(n, ra)))
	}
}

func main() {
	run(os.Stdout)
}
