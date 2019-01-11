// UVa 499 - What's The Frequency, Kenneth?

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func output(out *os.File, dict map[byte]int, max int) {
	var cs []int
	for k, v := range dict {
		if v == max {
			cs = append(cs, int(k))
		}
	}
	sort.Ints(cs)
	for _, v := range cs {
		fmt.Fprintf(out, "%c", v)
	}
	fmt.Fprintf(out, " %d\n", max)
}

func main() {
	in, _ := os.Open("499.in")
	defer in.Close()
	out, _ := os.Create("499.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		dict := make(map[byte]int)
		max, cnt := 0, 0
		for _, c := range []byte(s.Text()) {
			if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
				if cnt = dict[c] + 1; cnt > max {
					max = cnt
				}
				dict[c] = cnt
			}
		}
		output(out, dict, max)
	}
}
