package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var line string
	fmt.Fscan(reader, &line)

	anslist := make([]int, 26)
	for ind, value := range line {
		anslist[value-'a'] += (len(line) - ind) * (ind + 1)
	}
	for ind, value := range anslist {
		if value != 0 {
			fmt.Printf("%c: %d\n", rune(ind+int('a')), value)
		}
	}
}
