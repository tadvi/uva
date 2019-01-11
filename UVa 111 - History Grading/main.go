package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(s1, s2 []int) int {
	sz := len(s1)
	dp := make([][]int, sz+1)
	for i := range dp {
		dp[i] = make([]int, sz+1)
	}

	for i := 1; i <= sz; i++ {
		for j := 1; j <= sz; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[sz][sz]
}

func readLine(in io.Reader, fn func(s string)) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fn(scanner.Text())
	}
}

func getLine(rd io.Reader, n int) []int {
	var v int
	s1 := make([]int, n)
	for i := range s1 {
		fmt.Fscan(rd, &v)
		s1[v-1] = i + 1
	}
	return s1
}

func run(w io.Writer) {
	in, _ := os.Open("input.txt")
	defer in.Close()

	var n, line int
	var s1, s2 []int

	readLine(in, func(s string) {

		rd := strings.NewReader(s)

		if !strings.Contains(s, " ") {
			fmt.Fscan(rd, &n)
			line = 0
		} else {
			if line == 0 {
				s1 = getLine(rd, n)
				line++
			} else {
				s2 = getLine(rd, n)

				fmt.Println(lcs(s1, s2))
			}

		}
	})

}

func main() {
	run(os.Stdout)
}
