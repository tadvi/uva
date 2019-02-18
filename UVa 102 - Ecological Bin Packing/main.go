package main

import (
    "fmt"
    "io"
    "math"
    "os"
)

const (
    b = iota
    g
    c
)

type Rec struct {
    b, g, c int
}

type Result struct {
    sum    int
    best   int
    config string
}

func (res *Result) try(rec [3]Rec, conf string) {

    if cur := res.sum - rec[0].b - rec[1].g - rec[2].c; cur < res.best {
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
            best: math.MaxInt32,
            sum: rec[0].b + rec[1].b + rec[2].b +
                rec[0].g + rec[1].g + rec[2].g +
                rec[0].c + rec[1].c + rec[2].c,
        }

        res.try([3]Rec{rec[b], rec[c], rec[g]}, "BCG")
        res.try([3]Rec{rec[b], rec[g], rec[c]}, "BGC")
        res.try([3]Rec{rec[g], rec[c], rec[b]}, "CBG")
        res.try([3]Rec{rec[c], rec[g], rec[b]}, "CGB")
        res.try([3]Rec{rec[g], rec[b], rec[c]}, "GBC")
        res.try([3]Rec{rec[c], rec[b], rec[g]}, "GCB")
        fmt.Fprintf(w, "%s %d\n", res.config, res.best)
    }

}

func main() {
    run(os.Stdout)
}
