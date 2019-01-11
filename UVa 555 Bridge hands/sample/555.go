// UVa 555 - Bridge Hands

package main

import (
	"fmt"
	"os"
	"sort"
)

type card struct{ suit, number byte }

var (
	ptr          int
	hands        [4][]card
	directionMap = map[string]int{"S": 0, "W": 1, "N": 2, "E": 3}
	orders       = []string{"S", "W", "N", "E"}
	suitMap      = map[byte]int{'C': 0, 'D': 1, 'S': 2, 'H': 3}
	numberMap    = map[byte]int{'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7, 'T': 8, 'J': 9, 'Q': 10, 'K': 11, 'A': 12}
)

func place(line string) {
	n := len(line) / 2
	for i := 0; i < n; i++ {
		ptr = (ptr + 1) % 4
		hands[ptr] = append(hands[ptr], card{line[i*2], line[i*2+1]})
	}
}

func main() {
	in, _ := os.Open("555.in")
	defer in.Close()
	out, _ := os.Create("555.out")
	defer out.Close()

	var direction, line string
	for {
		if fmt.Fscanf(in, "%s", &direction); direction == "#" {
			break
		}
		ptr = directionMap[direction]
		for i := range hands {
			hands[i] = nil
		}
		for i := 0; i < 2; i++ {
			fmt.Fscanf(in, "%s", &line)
			place(line)
		}
		for i, order := range orders {
			fmt.Fprintf(out, "%s:", order)
			sort.Slice(hands[i], func(j, k int) bool {
				if suitMap[hands[i][j].suit] == suitMap[hands[i][k].suit] {
					return numberMap[hands[i][j].number] < numberMap[hands[i][k].number]
				}
				return suitMap[hands[i][j].suit] < suitMap[hands[i][k].suit]
			})
			for _, c := range hands[i] {
				fmt.Fprintf(out, " %c%c", c.suit, c.number)
			}
			fmt.Fprintln(out)
		}
	}
}
