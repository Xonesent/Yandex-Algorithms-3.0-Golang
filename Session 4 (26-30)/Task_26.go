package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(reader, &n, &m)

	matrix1 := make([][]int, n)
	matrix2 := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix1[i] = make([]int, m)
		matrix2[i] = make([]int, m)
		for ind := 0; ind < m; ind++ {
			fmt.Fscan(reader, &matrix1[i][ind])
		}
	}

	for i := 0; i < n; i++ {
		for ind := 0; ind < m; ind++ {
			if i == 0 && ind == 0 {
				matrix2[i][ind] = matrix1[i][ind]
			} else if i == 0 {
				matrix2[i][ind] = matrix2[i][ind-1] + matrix1[i][ind]
			} else if ind == 0 {
				matrix2[i][ind] = matrix2[i-1][ind] + matrix1[i][ind]
			} else {
				matrix2[i][ind] = min(matrix2[i][ind-1], matrix2[i-1][ind]) + matrix1[i][ind]
			}
		}
	}

	fmt.Println(matrix2[n-1][m-1])
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
