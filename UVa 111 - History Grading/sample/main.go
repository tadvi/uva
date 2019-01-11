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

func run(w io.Writer) {
	in, _ := os.Open("input.txt")
	defer in.Close()

	var n, v int
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		str := scanner.Text()
		rd := strings.NewReader(str)

		if !strings.Contains(str, " ") {
			fmt.Fscan(rd, &n)
		} else {
			s1 := make([]int, n)
			for i := range s1 {
				fmt.Fscan(rd, &v)
				s1[v-1] = i + 1
			}

			fmt.Println(s1)
		}

	}

	/*
		var n, v int
		if _, err := fmt.Fscan(in, &n); err != nil {

		}

		for {
			s1 := make([]int, n)
			for i := range s1 {
				fmt.Fscan(in, &v)
				s1[v-1] = i + 1
			}
			fmt.Println(s1)

			s2 := make([]int, n)

		sub:
			for {
				for i := 1; i <= n; i++ {
					if _, err := fmt.Fscan(in, &v); err != nil {
						break sub
					}
					s2[v-1] = i
				}
				fmt.Println(lcs(s1, s2))
			}

			n = v
			fmt.Println("new n", n)

		}*/
}

func main() {
	run(os.Stdout)
}
