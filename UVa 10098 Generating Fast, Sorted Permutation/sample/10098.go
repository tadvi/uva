// UVa 10098 - Generating Fast

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	res     map[string]bool
	out     *os.File
	visited map[int]bool
)

func dfs(strs []string, ans []string) {
	if len(ans) == len(strs) {
		s := strings.Join(ans, "")
		if _, ok := res[s]; !ok {
			res[s] = true
			fmt.Fprintln(out, s)
		}
		return
	}
	for i := range strs {
		if !visited[i] {
			visited[i] = true
			dfs(strs, append(ans, strs[i]))
			visited[i] = false
		}
	}
}

func main() {
	in, _ := os.Open("10098.in")
	defer in.Close()
	out, _ = os.Create("10098.out")
	defer out.Close()

	var n int
	var str string
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%s", &str)
		strs := make([]string, len(str))
		for i := range str {
			strs[i] = string(str[i])
		}
		sort.Strings(strs)
		res = make(map[string]bool)
		visited = make(map[int]bool)
		dfs(strs, nil)
		fmt.Fprintln(out)
	}
}
