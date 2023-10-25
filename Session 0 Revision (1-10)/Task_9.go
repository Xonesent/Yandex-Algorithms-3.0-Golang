package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	matrix1 := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix1[i] = make([]int, m)
		for ind := 0; ind < m; ind++ {
			fmt.Fscan(reader, &matrix1[i][ind])
		}
	}

	matrix2 := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		matrix2[i] = make([]int, m+1)
		for ind := 0; ind < m+1; ind++ {
			if i >= 1 && ind >= 1 {
				matrix2[i][ind] = matrix2[i][ind-1] + matrix2[i-1][ind] - matrix2[i-1][ind-1] + matrix1[i-1][ind-1]
			}
		}
	}

	for i := 0; i < k; i++ {
		var x1, y1, x2, y2 int
		fmt.Fscan(reader, &x1, &y1, &x2, &y2)
		fmt.Println(matrix2[x2][y2] - matrix2[x1-1][y2] - matrix2[x2][y1-1] + matrix2[x1-1][y1-1])
	}
}
