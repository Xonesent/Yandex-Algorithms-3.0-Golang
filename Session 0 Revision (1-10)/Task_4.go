package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k, row, place int
	fmt.Fscan(reader, &n, &k, &row, &place)

	place = (row-1)*2 + place
	row1 := (place + k + 1) / 2
	row2 := (place - k + 1) / 2

	switch {
	case place+k <= n && place+k > 0 && place-k <= n && place-k > 0:
		if row-row1 < row2-row {
			fmt.Println(row2, 2-(place-k)%2)
		} else {
			fmt.Println(row1, 2-(place+k)%2)
		}
	case place+k <= n && place+k > 0:
		fmt.Println(row1, 2-(place+k)%2)
	case place-k <= n && place-k > 0:
		fmt.Println(row2, 2-(place-k)%2)
	default:
		fmt.Println(-1)
	}
}
