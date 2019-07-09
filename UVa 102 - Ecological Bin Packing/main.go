package main

import (
	"fmt"
	"os"

	"github.com/tadvi/uva/functool"
)

func run() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	bins := []string{"B", "G", "C"}
	for {
		var row [9]int
		for i := 0; i < 9; i++ {
			if _, err := fmt.Fscan(f, &row[i]); err != nil {
				return
			}
		}

		min := 1 << 30
		var s string

		functool.Permutate([]int{0, 1, 2}, func(p []int) {
			total := 0
			for i := range bins {
				for j := range bins {
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
				s, min = name, total
			} else if total == min {
				if name < s {
					s = name
				}
			}
		})

		fmt.Println(s, min)
	}
}

func main() {
	run()
}
