package main

import (
	"fmt"
	"io"
	"os"
)

type Point struct {
	x, y float64
}

type Line struct {
	p1, p2 Point
}

func (ln Line) intersect(other Line) bool {
	return (ccw(ln.p1, ln.p2, other.p1)*ccw(ln.p1, ln.p2, other.p2) <= 0) &&
		(ccw(other.p1, other.p2, ln.p1)*ccw(other.p1, other.p2, ln.p2) <= 0)
}

// CCW (Counter Clockwise) Test
func ccw(p, q, r Point) int {
	if turn(p, q, r) > 0 {
		return 1
	}
	return 0
}

func turn(p, q, r Point) int {
	result := (r.x-q.x)*(p.y-q.y) - (r.y-q.y)*(p.x-q.x)
	if result < 0 {
		return -1
	}
	if result > 0 {
		return 1
	}
	return 0
}

func grahamSearch(polygon []*Point) []*Point {

	return nil
}

// TODO: not finished!!

func run(w io.Writer) {
	in, _ := os.Open("input.txt")
	defer in.Close()

	var hulls [][]*Point

	var n int
	var x, y float64
	for {
		var pts []*Point
		if _, err := fmt.Fscan(in, &n); err != nil || n == -1 {
			break
		}
		for i := 0; i < n; i++ {
			if _, err := fmt.Fscan(in, &x, &y); err != nil {
				panic("unable to read x, y")
			}
			pts = append(pts, &Point(x, y))
		}
		h := grahamSearch(pts)
		hulls = append(hulls, h)
	}

	// occupiedHulls = make(map[int]bool)
}

func main() {
	run(os.Stdout)
}
