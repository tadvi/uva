// UVa 825 - Walking on the Safe Side

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	w, n int
	grid [][]bool
)

func solve() int {
	dp := make([][]int, w+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][1] = 1
	for i := 1; i <= w; i++ {
		for j := 1; j <= n; j++ {
			if !grid[i-1][j-1] {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[w][n]
}

func main() {
	in, _ := os.Open("825.in")
	defer in.Close()
	out, _ := os.Create("825.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var kase, t int
	s.Scan()
	for fmt.Sscanf(s.Text(), "%d", &kase); kase > 0 && s.Scan(); kase-- {
		s.Scan()
		fmt.Sscanf(s.Text(), "%d%d", &w, &n)
		grid = make([][]bool, w)
		for i := range grid {
			grid[i] = make([]bool, n)
			s.Scan()
			for _, token := range strings.Fields(s.Text())[1:] {
				fmt.Sscanf(token, "%d", &t)
				grid[i][t-1] = true
			}
		}
		fmt.Fprintln(out, solve())
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
