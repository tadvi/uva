// UVa 384 - Slurpys

package main

import (
	"fmt"
	"os"
	"strings"
)

func isSlump(s string) int {
	if !strings.HasPrefix(s, "D") && !strings.HasPrefix(s, "E") || !strings.HasPrefix(s[1:], "F") {
		return 0
	}
	idx := 2
	for s[idx] == 'F' {
		idx++
	}
	if s[idx] == 'G' {
		return idx + 1
	}
	if len1 := isSlump(s[idx:]); len1 != 0 {
		return idx + len1
	}
	return 0
}

func isSlimp(s string) int {
	if strings.HasPrefix(s, "AH") {
		return 2
	}
	if strings.HasPrefix(s, "AB") {
		if len1 := isSlimp(s[2:]); len1 != 0 && s[2+len1] == 'C' {
			return 3 + len1
		}
		return 0
	}
	if len1 := isSlump(s[1:]); len1 != 0 && s[1+len1] == 'C' {
		return 2 + len1
	}
	return 0
}

func isSlurpy(s string) bool {
	if len1 := isSlimp(s); len1 != 0 {
		if len2 := isSlump(s[len1:]); len2 != 0 {
			return len1+len2 == len(s)
		}
		return false
	}
	return false
}

func main() {
	in, _ := os.Open("384.in")
	defer in.Close()
	out, _ := os.Create("384.out")
	defer out.Close()

	var kase int
	var line string
	fmt.Fprintln(out, "SLURPYS OUTPUT")
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		if fmt.Fscanf(in, "%s", &line); isSlurpy(line) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
	fmt.Fprintln(out, "END OF OUTPUT")
}
