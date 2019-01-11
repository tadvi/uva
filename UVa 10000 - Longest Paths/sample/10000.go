// UVa 10000 - Longest Paths

package main

import (
	"fmt"
	"math"
	"os"
)

var (
	lowest, longest int
	distance        []int
	adj             [][]bool
)

func dfs(s int) {
	for i := range adj[s] {
		if adj[s][i] {
			if distance[s]+1 > distance[i] {
				switch distance[i] = distance[s] + 1; {
				case distance[i] > longest:
					longest = distance[i]
					lowest = i
				case distance[i] == longest && i < lowest:
					lowest = i
				}
			}
			dfs(i)
		}
	}
}

func main() {
	in, _ := os.Open("10000.in")
	defer in.Close()
	out, _ := os.Create("10000.out")
	defer out.Close()

	var n, s, n1, n2 int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fscanf(in, "%d", &s)
		adj = make([][]bool, n+1)
		for i := range adj {
			adj[i] = make([]bool, n+1)
		}
		for {
			if fmt.Fscanf(in, "%d%d", &n1, &n2); n1 == 0 && n2 == 0 {
				break
			}
			adj[n1][n2] = true
		}
		distance = make([]int, n+1)
		longest, lowest = 0, math.MaxInt32
		dfs(s)
		fmt.Fprintf(out, "Case %d: The longest path from %d has length %d, finishing at %d\n\n", kase, s, longest, lowest)
	}
}
