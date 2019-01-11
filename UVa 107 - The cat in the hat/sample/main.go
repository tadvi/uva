package main

import (
	"fmt"
	"io"
	"os"
)

func run(w io.Writer) {
	in, _ := os.Open("input.txt")
	defer in.Close()

	var n int
	for {
		if _, err := fmt.Fscan(in, &n); err != nil {
			break
		}

		// TODO: Add your code here

		var w int
	}
}

func main() {
	run(os.Stdout)
}
