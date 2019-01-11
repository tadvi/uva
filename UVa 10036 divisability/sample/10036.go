// UVa 10036 - Divisibility

package main

import (
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func divisible(n, k int, num []int) bool {
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, k)
	}
	dp[0][num[0]%k] = true
	for i := 1; i < n; i++ {
		for j := range dp[i-1] {
			if dp[i-1][j] {
				dp[i][(j+num[i])%k] = true
				dp[i][(j-num[i]+k)%k] = true
			}
		}
	}
	return dp[n-1][0]
}

func main() {
	in, _ := os.Open("10036.in")
	defer in.Close()
	out, _ := os.Create("10036.out")
	defer out.Close()

	var kase, n, k, tmp int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d", &n, &k)
		num := make([]int, n)
		for i := range num {
			fmt.Fscanf(in, "%d", &tmp)
			num[i] = abs(tmp) % k
		}
		if divisible(n, k, num) {
			fmt.Fprintln(out, "Divisible")
		} else {
			fmt.Fprintln(out, "Not divisible")
		}
	}
}
