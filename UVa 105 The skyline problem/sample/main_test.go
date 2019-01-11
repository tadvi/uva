package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	want := []string{
		"1 11 3 13 9 0 12 7 16 3 19 18 22 3 23 13 29 0\n",
	}

	var b bytes.Buffer
	run(&b)

	for i := 0; ; i++ {
		s, err := b.ReadString('\n')
		if err != nil {
			break
		}

		if s != want[i] {
			t.Errorf("got %q; want %q", s, want[i])
		}

	}
}
