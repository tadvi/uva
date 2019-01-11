package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type Rec struct {
	b, g, c int
}

type Result struct {
	blueSum, greenSum, clearSum int

	best   int
	config string
}

func (res *Result) try(rec [3]Rec, conf string) {
	cur := (res.blueSum - rec[0].b) + (res.greenSum - rec[1].g) + (res.clearSum - rec[2].c)
	if cur < res.best {
		res.best = cur
		res.config = conf
	}
}

func run(w io.Writer) {
	r, _ := os.Open("input.txt")
	defer r.Close()

	for {
		var rec [3]Rec

		for i := 0; i < 3; i++ {
			if _, err := fmt.Fscan(r, &rec[i].b, &rec[i].g, &rec[i].c); err == io.EOF {
				return
			}
		}

		res := Result{
			best:     math.MaxInt32,
			blueSum:  rec[0].b + rec[1].b + rec[2].b,
			greenSum: rec[0].g + rec[1].g + rec[2].g,
			clearSum: rec[0].c + rec[1].c + rec[2].c,
		}

		res.try([3]Rec{rec[0], rec[2], rec[1]}, "BCG")
		res.try([3]Rec{rec[0], rec[1], rec[2]}, "BGC")
		res.try([3]Rec{rec[1], rec[2], rec[0]}, "CBG")
		res.try([3]Rec{rec[2], rec[1], rec[0]}, "CGB")
		res.try([3]Rec{rec[1], rec[0], rec[2]}, "GBC")
		res.try([3]Rec{rec[2], rec[0], rec[1]}, "GCB")

		fmt.Fprintf(w, "%s %d\n", res.config, res.best)
	}
}

func main() {
	run(os.Stdout)
}
