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

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}

	matrix[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i+1 < n && j+2 < m {
				matrix[i+1][j+2] += matrix[i][j]
			}
			if i+2 < n && j+1 < m {
				matrix[i+2][j+1] += matrix[i][j]
			}
		}
	}

	fmt.Println(matrix[n-1][m-1])
}
