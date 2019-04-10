package main

import (
	"fmt"
	"io"
	"os"
)

func perm(a []int, start int, fn func([]int)) {
	if start == len(a)-1 {
		fn(a)
		return
	}
	for i := start; i < len(a); i++ {
		a[i], a[start] = a[start], a[i]
		perm(a, start+1, fn)
		a[i], a[start] = a[start], a[i]
	}
}

func run(w io.Writer) {
	r, _ := os.Open("input.txt")
	bins := []string{"B", "G", "C"}

	for {

		var row [9]int
		for i := 0; i < 9; i++ {
			if _, err := fmt.Fscan(r, &row[i]); err != nil {
				return
			}
		}

		min := 1 << 30
		var curr string

		perm([]int{0, 1, 2}, 0, func(p []int) {
			total := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if j != p[i] {
						total += row[i*3+j]
					}
				}
			}

			var name string
			for _, i := range p {
				name += bins[i]
			}

			if total < min {
				curr, min = name, total
			} else if total == min {
				if name < curr {
					curr = name
				}
			}
		})

		fmt.Println(curr, min)
	}

}

func main() {
	run(os.Stdout)
}
