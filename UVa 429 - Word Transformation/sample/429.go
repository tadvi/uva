// UVa 429 - Word Transformation

package main

import (
	"fmt"
	"os"
)

func sameLenAndOneCharDiff(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}
	count := 0
	for i := range w1 {
		if w1[i] != w2[i] {
			count++
		}
	}
	return count == 1
}

func buildLink(dict map[string][]string, word string) {
	dict[word] = nil
	for k, v := range dict {
		if k != word && sameLenAndOneCharDiff(k, word) {
			dict[k] = append(v, word)
			dict[word] = append(dict[word], k)
		}
	}
}

type node struct {
	w string
	n int
}

func bfs(dict map[string][]string, fm, to string) int {
	for visited, queue := map[string]bool{fm: true}, []node{{fm, 0}}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		adjs := dict[curr.w]
		for _, v := range adjs {
			if v == to {
				return curr.n + 1
			}
			if _, ok := visited[v]; !ok {
				visited[v] = true
				queue = append(queue, node{v, curr.n + 1})
			}
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("429.in")
	defer in.Close()
	out, _ := os.Create("429.out")
	defer out.Close()

	var n int
	var word, fm, to string
	for fmt.Fscanf(in, "%d\n\n", &n); n > 0; n-- {
		dict := make(map[string][]string)
		for {
			if fmt.Fscanf(in, "%s", &word); word == "*" {
				break
			}
			buildLink(dict, word)
		}
		for {
			if _, err := fmt.Fscanf(in, "%s%s\n", &fm, &to); err != nil {
				break
			}
			fmt.Fprintln(out, fm, to, bfs(dict, fm, to))
		}
	}
}
