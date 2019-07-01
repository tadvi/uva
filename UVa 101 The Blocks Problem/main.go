// UVa 101 - The Blocks Problem
package main

import (
	"fmt"
	"os"
)

// Matrix holds all blocks.
type Matrix struct {
	blocks [][]int
}

func newBlocks(n int) [][]int {
	b := make([][]int, n)
	for i := range b {
		b[i] = []int{i}
	}
	return b
}

func (mat *Matrix) find(n int) int {
	for i := range mat.blocks {
		for j := range mat.blocks[i] {
			if mat.blocks[i][j] == n {
				return i
			}
		}
	}
	panic("not found")
}

func (mat *Matrix) reset(n int) {
	stack := mat.blocks[mat.find(n)]

	for last := len(stack) - 1; stack[last] != n; last = len(stack) - 1 {
		i := stack[last]
		mat.blocks[i] = append(mat.blocks[i], i)
		stack = stack[:last]
	}
}

func (mat *Matrix) move(a, b int) {
	p1, p2 := mat.find(a), mat.find(b)
	mat.blocks[p1] = mat.blocks[p1][:len(mat.blocks[p1])-1]
	mat.blocks[p2] = append(mat.blocks[p2], a)
}

func (mat *Matrix) pile(a, b int) {
	p1, p2 := mat.find(a), mat.find(b)
	for i, val := range mat.blocks[p1] {
		if val == a {
			blocksToMove := mat.blocks[p1][i:]
			mat.blocks[p1] = mat.blocks[p1][:i]
			mat.blocks[p2] = append(mat.blocks[p2], blocksToMove...)
			return
		}
	}
}

func (mat *Matrix) process(act, pos string, a, b int) {
	if act == "move" {
		mat.reset(a)
	}
	if pos == "onto" {
		mat.reset(b)
	}

	switch act {
	case "move":
		mat.move(a, b)
	case "pile":
		mat.pile(a, b)
	}
}

func run() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var n, a, b int
	var act, pos string
	fmt.Fscan(f, &n)

	mat := &Matrix{blocks: newBlocks(n)}

	for {
		if fmt.Fscan(f, &act); act == "quit" {
			break
		}
		fmt.Fscan(f, &a, &pos, &b)
		mat.process(act, pos, a, b)
	}

	for i, line := range mat.blocks {
		fmt.Printf("%d:", i)
		for _, val := range line {
			fmt.Printf(" %d", val)
		}
		fmt.Println()
	}
}

func main() {
	run()
}
