package main

import (
	"fmt"
	"io"
	"os"
)

func run(w io.Writer) {
	in, _ := os.Open("input.txt")
	defer in.Close()

	var l, h, r, maxr int
	var city [10001]int
	for {
		if _, err := fmt.Fscan(in, &l, &h, &r); err != nil {
			break
		}

		for i := l; i < r; i++ {
			if city[i] < h {
				city[i] = h
			}
		}

		var w int

		if r > maxr {
			maxr = r
		}
	}

	for i := 1; i < maxr; i++ {
		if city[i-1] != city[i] {
			fmt.Printf("%d %d ", i, city[i])
		}
	}
	fmt.Println(maxr, 0)
}

func main() {
	run(os.Stdout)
}
