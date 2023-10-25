package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, s, t, q int
	fmt.Fscan(reader, &n, &m, &s, &t, &q)

	sum := 0
	matrix := create_filling_matrix(n, m, s, t)
	for i := 0; i < q; i++ {
		var value1, value2 int
		fmt.Fscan(reader, &value1, &value2)
		if matrix[value1-1][value2-1] == -1 {
			sum = -1
			break
		} else {
			sum += matrix[value1-1][value2-1]
		}
	}
	fmt.Print(sum)
}

func create_filling_matrix(n, m, s, t int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for ind := 0; ind < m; ind++ {
			matrix[i][ind] = -1
		}
	}
	matrix[s-1][t-1] = 0

	queue := make(map[int][][2]int)
	wave := 0
	queue[wave] = append(queue[wave], [2]int{s - 1, t - 1})

	for wave <= n*m {
		for len(queue[wave]) > 0 {
			curr := queue[wave][0]
			queue[wave] = queue[wave][1:]

			if curr[0]-2 >= 0 && curr[1]-1 >= 0 && matrix[curr[0]-2][curr[1]-1] == -1 {
				matrix[curr[0]-2][curr[1]-1] = wave + 1
				queue[wave+1] = append(queue[wave+1], [2]int{curr[0] - 2, curr[1] - 1})
			}
			if curr[0]-1 >= 0 && curr[1]-2 >= 0 && matrix[curr[0]-1][curr[1]-2] == -1 {
				matrix[curr[0]-1][curr[1]-2] = wave + 1
				queue[wave+1] = append(queue[wave+1], [2]int{curr[0] - 1, curr[1] - 2})
			}
			if curr[0]-2 >= 0 && curr[1]+1 < m && matrix[curr[0]-2][curr[1]+1] == -1 {
				matrix[curr[0]-2][curr[1]+1] = wave + 1
				queue[wave+1] = append(queue[wave+1], [2]int{curr[0] - 2, curr[1] + 1})
			}
			if curr[0]-1 >= 0 && curr[1]+2 < m && matrix[curr[0]-1][curr[1]+2] == -1 {
				matrix[curr[0]-1][curr[1]+2] = wave + 1
				queue[wave+1] = append(queue[wave+1], [2]int{curr[0] - 1, curr[1] + 2})
			}
			if curr[0]+2 < n && curr[1]-1 >= 0 && matrix[curr[0]+2][curr[1]-1] == -1 {
				matrix[curr[0]+2][curr[1]-1] = wave + 1
				queue[wave+1] = append(queue[wave+1], [2]int{curr[0] + 2, curr[1] - 1})
			}
			if curr[0]+1 < n && curr[1]-2 >= 0 && matrix[curr[0]+1][curr[1]-2] == -1 {
				matrix[curr[0]+1][curr[1]-2] = wave + 1
				queue[wave+1] = append(queue[wave+1], [2]int{curr[0] + 1, curr[1] - 2})
			}
			if curr[0]+2 < n && curr[1]+1 < m && matrix[curr[0]+2][curr[1]+1] == -1 {
				matrix[curr[0]+2][curr[1]+1] = wave + 1
				queue[wave+1] = append(queue[wave+1], [2]int{curr[0] + 2, curr[1] + 1})
			}
			if curr[0]+1 < n && curr[1]+2 < m && matrix[curr[0]+1][curr[1]+2] == -1 {
				matrix[curr[0]+1][curr[1]+2] = wave + 1
				queue[wave+1] = append(queue[wave+1], [2]int{curr[0] + 1, curr[1] + 2})
			}
		}
		wave++
	}

	return matrix
}
