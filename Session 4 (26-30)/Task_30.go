package main

import (
	"bufio"
	"fmt"
	"os"
)

type ans struct {
	count   int
	anslist []int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)
	num1 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &num1[i])
	}

	var m int
	fmt.Fscan(reader, &m)
	num2 := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &num2[i])
	}

	matrix := make([][]ans, n+1)
	for i := 0; i < n+1; i++ {
		matrix[i] = make([]ans, m+1)
		for ind := 0; ind < m+1; ind++ {
			matrix[i][ind].count = 0
		}
	}

	matrix_filling(matrix, n, m, num1, num2)
	for i := 0; i < len(matrix[n][m].anslist); i++ {
		fmt.Print(matrix[n][m].anslist[i], " ")
	}

	// for i := 0; i < n+1; i++ {
	// 	for ind := 0; ind < m+1; ind++ {
	// 		fmt.Print(matrix[i][ind].count, " ")
	// 	}
	// 	fmt.Println()
	// }
}

func matrix_filling(matrix [][]ans, n, m int, num1, num2 []int) {
	for i := 1; i < n+1; i++ {
		for ind := 1; ind < m+1; ind++ {
			if num1[i-1] == num2[ind-1] {
				matrix[i][ind] = matrix[i-1][ind-1]
				matrix[i][ind].count += 1
				matrix[i][ind].anslist = append(matrix[i][ind].anslist, num1[i-1])
			} else {
				matrix[i][ind] = max(matrix[i-1][ind], matrix[i][ind-1])
			}
		}
	}
}

func max(a ans, b ans) ans {
	if a.count > b.count {
		return a
	} else {
		return b
	}
}
