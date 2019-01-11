package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	// TODO: Add your test cases.
	want := []string{
		"1 2 1\n",
		"1 2 4 1\n",
		"no arbitrage sequence exists\n",
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
