// UVa 101 - The Blocks Problem
package main

import (
	"fmt"
	"io"
	"os"
)

type Matrix struct {
	blocks [][]int
}

func (mat *Matrix) initBlocks(n int) {
	mat.blocks = make([][]int, n)
	for i := range mat.blocks {
		mat.blocks[i] = []int{i}
	}
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
		mat.blocks[stack[last]] = append(mat.blocks[stack[last]], stack[last])
		stack = stack[:last]
	}
}

func (mat *Matrix) move(n1, n2 int) {
	pos1, pos2 := mat.find(n1), mat.find(n2)
	mat.blocks[pos1] = mat.blocks[pos1][:len(mat.blocks[pos1])-1]
	mat.blocks[pos2] = append(mat.blocks[pos2], n1)
}

func (mat *Matrix) pile(n1, n2 int) {
	pos1, pos2 := mat.find(n1), mat.find(n2)
	for i, val := range mat.blocks[pos1] {
		if val == n1 {
			blocksToMove := mat.blocks[pos1][i:]
			mat.blocks[pos1] = mat.blocks[pos1][:i]
			mat.blocks[pos2] = append(mat.blocks[pos2], blocksToMove...)
			return
		}
	}
}

func (mat *Matrix) process(s1, s2 string, n1, n2 int) {
	if s1 == "move" {
		mat.reset(n1)
	}
	if s2 == "onto" {
		mat.reset(n2)
	}

	switch s1 {
	case "move":
		mat.move(n1, n2)
	case "pile":
		mat.pile(n1, n2)
	}
}

func run(w io.Writer) {
	in, _ := os.Open("input.txt")
	defer in.Close()

	var n, n1, n2 int
	var s1, s2 string
	fmt.Fscan(in, &n)

	mat := &Matrix{}
	mat.initBlocks(n)
	for {
		if fmt.Fscan(in, &s1); s1 == "quit" {
			break
		}
		fmt.Fscan(in, &n1, &s2, &n2)
		mat.process(s1, s2, n1, n2)
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
	run(os.Stdout)
}
