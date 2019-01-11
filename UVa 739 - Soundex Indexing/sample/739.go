// UVa 739 - Soundex Indexing

package main

import (
	"fmt"
	"os"
	"strings"
)

var codeMap = map[byte]string{
	'B': "1", 'P': "1", 'F': "1", 'V': "1",
	'C': "2", 'S': "2", 'K': "2", 'G': "2", 'J': "2", 'Q': "2", 'X': "2", 'Z': "2",
	'D': "3", 'T': "3",
	'L': "4",
	'M': "5", 'N': "5",
	'R': "6",
}

func encode(name string) string {
	skipped := false
	soundex := name[0:1]
	for i := 1; i < len(name); i++ {
		if code := codeMap[name[i]]; code != "" && code != soundex[len(soundex)-1:] {
			if len(soundex) == 1 && code == codeMap[name[0]] && !skipped {
				skipped = true
				continue
			}
			if soundex += code; len(soundex) == 4 {
				break
			}
		}
	}
	if len(soundex) < 4 {
		soundex += strings.Repeat("0", 4-len(soundex))
	}
	return soundex
}

func main() {
	in, _ := os.Open("739.in")
	defer in.Close()
	out, _ := os.Create("739.out")
	defer out.Close()

	fmt.Fprintln(out, "         NAME                     SOUNDEX CODE")

	var name string
	for {
		if _, err := fmt.Fscanf(in, "%s", &name); err != nil {
			break
		}
		encoded := encode(name)
		if len(name) < 25 {
			name += strings.Repeat(" ", 25-len(name))
		}
		fmt.Fprintf(out, "         %s%s\n", name, encoded)
	}
	fmt.Fprintln(out, "                   END OF OUTPUT")
}
