package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	graph := make(map[int][]int)
	matrix_scan(graph, reader)

	var start, end int
	fmt.Fscan(reader, &start, &end)

	if len(graph[start]) == 0 || len(graph[end]) == 0 {
		fmt.Println(-1)
	} else if start == end {
		fmt.Println(0)
	} else {
		fmt.Println(findShortcut(graph, start, end))
	}
}

func findShortcut(graph map[int][]int, start, end int) int {
	visited := make(map[int]bool)
	queue := make(map[int][]int)
	wave := 0

	visited[start] = true
	queue[wave] = append(queue[wave], start)

	for wave <= len(graph) {
		for len(queue[wave]) > 0 {
			parent := queue[wave][0]
			queue[wave] = queue[wave][1:]

			for _, child := range graph[parent] {
				if !visited[child] {
					if child == end {
						return wave + 1
					}
					visited[child] = true
					queue[wave+1] = append(queue[wave+1], child)
				}
			}
		}
		wave++
	}

	return -1
}

func matrix_scan(graph map[int][]int, reader *bufio.Reader) {
	var n, curr_value int
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		for ind := 0; ind < n; ind++ {
			fmt.Fscan(reader, &curr_value)
			if curr_value == 1 {
				graph[i+1] = append(graph[i+1], ind+1)
			}
		}
	}
}
