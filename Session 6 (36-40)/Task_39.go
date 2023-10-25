package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	matrix, position := create_matrix()
	fmt.Println(find_path(matrix, position))
}

func find_path(matrix [][][]byte, position [3]int) int {
	if position[0] == 0 {
		return 0
	}

	queue := make(map[int][][3]int)
	wave := 0
	n := len(matrix)

	queue[wave] = append(queue[wave], position)

	for wave <= n*n*n {
		for len(queue[wave]) > 0 {
			start := queue[wave][0]
			queue[wave] = queue[wave][1:]

			if start[1]-1 >= 0 && matrix[start[0]][start[1]-1][start[2]] != '#' && matrix[start[0]][start[1]-1][start[2]] == '.' {
				matrix[start[0]][start[1]-1][start[2]] = byte(wave) + 1
				queue[wave+1] = append(queue[wave+1], [3]int{start[0], start[1] - 1, start[2]})
			}
			if start[1]+1 < n && matrix[start[0]][start[1]+1][start[2]] != '#' && matrix[start[0]][start[1]+1][start[2]] == '.' {
				matrix[start[0]][start[1]+1][start[2]] = byte(wave) + 1
				queue[wave+1] = append(queue[wave+1], [3]int{start[0], start[1] + 1, start[2]})
			}
			if start[2]-1 >= 0 && matrix[start[0]][start[1]][start[2]-1] != '#' && matrix[start[0]][start[1]][start[2]-1] == '.' {
				matrix[start[0]][start[1]][start[2]-1] = byte(wave) + 1
				queue[wave+1] = append(queue[wave+1], [3]int{start[0], start[1], start[2] - 1})
			}
			if start[2]+1 < n && matrix[start[0]][start[1]][start[2]+1] != '#' && matrix[start[0]][start[1]][start[2]+1] == '.' {
				matrix[start[0]][start[1]][start[2]+1] = byte(wave) + 1
				queue[wave+1] = append(queue[wave+1], [3]int{start[0], start[1], start[2] + 1})
			}
			if start[0]+1 < n && matrix[start[0]+1][start[1]][start[2]] != '#' && matrix[start[0]+1][start[1]][start[2]] == '.' {
				matrix[start[0]+1][start[1]][start[2]] = byte(wave) + 1
				queue[wave+1] = append(queue[wave+1], [3]int{start[0] + 1, start[1], start[2]})
			}
			if start[0]-1 >= 0 && matrix[start[0]-1][start[1]][start[2]] != '#' && matrix[start[0]-1][start[1]][start[2]] == '.' {
				if start[0]-1 == 0 {
					return wave + 1
				}
				matrix[start[0]-1][start[1]][start[2]] = byte(wave) + 1
				queue[wave+1] = append(queue[wave+1], [3]int{start[0] - 1, start[1], start[2]})
			}
		}

		wave++
	}

	return -1
}

func create_matrix() ([][][]byte, [3]int) {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var pos, return_value [3]int
	fmt.Fscan(reader, &n)

	matrix := make([][][]byte, n)
	for i := 0; i < n; i++ {
		matrix[i], pos = scan_one_matrix(reader, n, i)
		if pos[2] != -1 {
			return_value = pos
		}
	}

	return matrix, return_value
}

func scan_one_matrix(reader *bufio.Reader, n int, index int) ([][]byte, [3]int) {
	matrix := make([][]byte, n)
	return_value := [3]int{index, -1, -1}

	place := 0
	startPos := 0

	for i := 0; i < n; i++ {
		var row string
		fmt.Scan(&row)
		matrix[i] = []byte(row)
		if strings.Contains(row, "S") {
			startPos = strings.Index(row, "S")
			return_value = [3]int{index, place, startPos}
		}
		place++
	}

	return matrix, return_value
}
