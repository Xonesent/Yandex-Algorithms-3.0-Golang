package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	map1 := make(map[rune]int)
	anslist := make([]int, 0)
	var max int = 0

	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimRight(line, "\n")
		if line == "" {
			break
		}
		for _, value := range line {
			if value != ' ' {
				map1[value] += 1
				if map1[value] > max {
					max = map1[value]
				}
				if map1[value] == 1 {
					anslist = append(anslist, int(value))
				}
			}
		}
	}

	sort.Ints(anslist)

	ansmatrix := make([][]string, max+1)
	for i := 0; i < max+1; i++ {
		ansmatrix[i] = make([]string, len(anslist))
		if i == max {
			for ind := 0; ind < len(anslist); ind++ {
				ansmatrix[i][ind] = string(rune(anslist[ind]))
			}
		} else {
			for ind := 0; ind < len(anslist); ind++ {
				if max-i <= map1[rune(anslist[ind])] {
					ansmatrix[i][ind] = "#"
				} else {
					ansmatrix[i][ind] = " "
				}
			}

		}
	}

	for i := 0; i < max+1; i++ {
		for ind := 0; ind < len(anslist); ind++ {
			fmt.Print(ansmatrix[i][ind])
		}
		fmt.Println()
	}
}
