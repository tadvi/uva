package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	want := []string{
		"5\n",
		"3 1 2 4 5\n",
		"4\n",
		"7 2 5 6\n",
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
