// UVa 11340 - Newspaper

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("11340.in")
	defer in.Close()
	out, _ := os.Create("11340.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var kase, t, k, l int
	var c byte
	s.Scan()
	for fmt.Sscanf(s.Text(), "%d", &kase); kase > 0 && s.Scan(); kase-- {
		paid := make(map[byte]int)
		for fmt.Sscanf(s.Text(), "%d", &k); k > 0 && s.Scan(); k-- {
			fmt.Sscanf(s.Text(), "%c%d", &c, &t)
			paid[c] = t
		}
		var total float64
		s.Scan()
		for fmt.Sscanf(s.Text(), "%d", &l); l > 0 && s.Scan(); l-- {
			line := s.Text()
			for i := range line {
				total += float64(paid[line[i]])
			}
		}
		fmt.Fprintf(out, "%.2f$\n", total/100)
	}
}
