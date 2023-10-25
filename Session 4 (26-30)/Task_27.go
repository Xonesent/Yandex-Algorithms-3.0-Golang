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

	scan_matrix2(matrix1, matrix2, n, m)
	anslist := find_path(matrix2, n, m)

	fmt.Println(matrix2[n-1][m-1])
	for i := len(anslist) - 1; i >= 0; i-- {
		fmt.Print(anslist[i], " ")
	}
}

func scan_matrix2(matrix1, matrix2 [][]int, n int, m int) {
	for i := 0; i < n; i++ {
		for ind := 0; ind < m; ind++ {
			if i == 0 && ind == 0 {
				matrix2[i][ind] = matrix1[i][ind]
			} else if i == 0 {
				matrix2[i][ind] = matrix2[i][ind-1] + matrix1[i][ind]
			} else if ind == 0 {
				matrix2[i][ind] = matrix2[i-1][ind] + matrix1[i][ind]
			} else {
				matrix2[i][ind] = max(matrix2[i][ind-1], matrix2[i-1][ind]) + matrix1[i][ind]
			}
		}
	}
}

func find_path(matrix2 [][]int, n int, m int) []string {
	var x, y int = m - 1, n - 1
	anslist := make([]string, 0)
	for x != 0 || y != 0 {
		if x-1 >= 0 && y-1 >= 0 {
			if matrix2[y][x-1] > matrix2[y-1][x] {
				anslist = append(anslist, "R")
				x--
			} else {
				anslist = append(anslist, "D")
				y--
			}
		} else if x-1 >= 0 {
			anslist = append(anslist, "R")
			x--
		} else if y-1 >= 0 {
			anslist = append(anslist, "D")
			y--
		}
	}

	return anslist
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
