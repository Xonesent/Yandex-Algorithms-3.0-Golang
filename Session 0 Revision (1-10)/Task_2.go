package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k int
	fmt.Fscanf(reader, "%d\n", &k)
	line1, _ := reader.ReadString('\n')
	line1 = strings.TrimRight(line1, "\n")

	var max int = 0
	for _, value := range "abcdefghijklmnopqrstuvwxyz" {
		right := 0
		bonus := k

		for left := 0; left < len(line1)-k; left++ {
			for right < len(line1) && bonus >= 0 {
				if rune(line1[right]) != value {
					bonus--
				}
				right++
			}

			if rune(line1[left]) != value {
				bonus++
			}

			if right-left-1 > max {
				max = right - left - 1
			}
		}
	}

	fmt.Println(max)
}
