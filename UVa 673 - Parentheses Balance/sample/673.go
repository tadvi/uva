// UVa 673 - Parentheses Balance

package main

import (
	"bufio"
	"fmt"
	"os"
)

func match(a, b byte) bool { return a == ')' && b == '(' || a == ']' && b == '[' }

func valid(line string) bool {
	var stack []byte
	for i := range line {
		switch line[i] {
		case '(', '[':
			stack = append(stack, line[i])
		case ')', ']':
			if len(stack) > 0 && match(line[i], stack[len(stack)-1]) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	in, _ := os.Open("673.in")
	defer in.Close()
	out, _ := os.Create("673.out")
	defer out.Close()

	var n int
	fmt.Fscanf(in, "%d", &n)

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)
	for ; n > 0 && s.Scan(); n-- {
		if valid(s.Text()) {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
