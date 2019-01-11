// UVa 10067 - Playing with Wheels

package main

import (
	"fmt"
	"os"
)

type config struct {
	n     [4]int
	steps int
}

var (
	forbidden  map[[4]int]bool
	target     config
	directions = [2]int{-1, 1}
)

func rotate(n [4]int, i, step int) [4]int {
	switch n[i] += step; n[i] {
	case -1:
		n[i] = 9
	case 10:
		n[i] = 0
	}
	return n
}

func bfs(initial config) int {
	for visited, queue := map[[4]int]bool{initial.n: true}, []config{initial}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		if curr.n == target.n {
			return curr.steps
		}
		for i := range curr.n {
			for _, direction := range directions {
				next := rotate(curr.n, i, direction)
				if !visited[next] && !forbidden[next] {
					visited[next] = true
					queue = append(queue, config{next, curr.steps + 1})
				}
			}
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("10067.in")
	defer in.Close()
	out, _ := os.Create("10067.out")
	defer out.Close()

	var kase, f int
	var t [4]int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d%d%d", &t[0], &t[1], &t[2], &t[3])
		initial := config{t, 0}
		fmt.Fscanf(in, "%d%d%d%d", &t[0], &t[1], &t[2], &t[3])
		target = config{t, -1}
		forbidden = make(map[[4]int]bool)
		for fmt.Fscanf(in, "%d", &f); f > 0; f-- {
			fmt.Fscanf(in, "%d%d%d%d", &t[0], &t[1], &t[2], &t[3])
			forbidden[t] = true
		}
		fmt.Fprintln(out, bfs(initial))
		if kase > 1 {
			fmt.Fscanln(in)
		}
	}
}
