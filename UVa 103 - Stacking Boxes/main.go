package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func isSame(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func fit(a, b []int) bool {
	for i := range a {
		if a[i] >= b[i] {
			return false
		}
	}
	return true
}

func order(n int, boxes [][]int) []int {
	sort.Slice(boxes, func(i, j int) bool {
		for idx := range boxes[i] {
			if boxes[i][idx] != boxes[j][idx] {
				return boxes[i][idx] < boxes[j][idx]
			}
		}
		return false
	})

	length := make([]int, n)
	for i := range length {
		length[i] = 1
	}
	pre := make([]int, n)
	for i := range pre {
		pre[i] = -1
	}

	max, maxIdx := 0, 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if fit(boxes[i], boxes[j]) {

				advance := length[i] + 1
				if advance > length[j] {
					if length[j] = advance; length[j] > max {
						max = length[j]
						maxIdx = j
					}
					pre[j] = i
				}

			}
		}
	}

	ret := []int{maxIdx}
	for pre[maxIdx] != -1 {
		maxIdx = pre[maxIdx]
		ret = append([]int{maxIdx}, ret...)
	}

	return ret
}

func output(ordered []int, boxes, original [][]int) {
	var ret []string
	for _, v := range ordered {
		for j := range original {
			if isSame(boxes[v], original[j]) {
				ret = append(ret, strconv.Itoa(j+1))
				break
			}
		}
	}
	fmt.Println(len(ret))
	fmt.Println(strings.Join(ret, " "))
}

func run(w io.Writer) {
	r, _ := os.Open("input.txt")
	defer r.Close()

	var n, d int

	for {
		if _, err := fmt.Fscan(r, &n, &d); err != nil {
			break
		}

		boxes := make([][]int, n)
		for i := range boxes {
			boxes[i] = make([]int, d)
			for j := range boxes[i] {
				fmt.Fscan(r, &boxes[i][j])
			}
			sort.Ints(boxes[i])
		}

		original := make([][]int, n)
		copy(original, boxes)

		ord := order(n, boxes)
		output(ord, boxes, original)
	}
}

func main() {
	run(os.Stdout)
}
