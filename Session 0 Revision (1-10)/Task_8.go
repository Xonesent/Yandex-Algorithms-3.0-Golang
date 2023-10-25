package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	var x1, y1, x2, y2 int = math.MaxInt64, math.MaxInt64, math.MinInt64, math.MinInt64
	for i := 0; i < n; i++ {
		var left, right int
		fmt.Fscan(reader, &left, &right)
		x1 = min(x1, left)
		x2 = max(x2, left)
		y1 = min(y1, right)
		y2 = max(y2, right)
	}

	fmt.Println(x1, y1, x2, y2)
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
