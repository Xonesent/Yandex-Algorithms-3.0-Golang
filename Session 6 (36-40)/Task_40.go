package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	graph := create_graph(reader)

	var start, end int
	fmt.Fscan(reader, &start, &end)

	fmt.Println(find_min_path(graph, start, end))
}

func find_min_path(graph map[int][]int, start, end int) int {
	if len(graph[start]) == 0 || len(graph[end]) == 0 {
		fmt.Println(-1)
	} else if start == end {
		return 0
	}
	visited := make(map[int]bool)
	queue := make(map[int][]int)
	wave := 0

	queue[wave] = append(queue[wave], start)
	visited[start] = true

	for wave <= len(graph) {
		for len(queue[wave]) > 0 {
			parent := queue[wave][0]
			queue[wave] = queue[wave][1:]

			for _, child := range graph[parent] {
				if !visited[child] {
					visited[child] = true
					queue[wave+1] = append(queue[wave+1], child)

					if child == end {
						return wave
					}
				}
			}
		}
		wave++
	}

	return -1
}

func create_graph(reader *bufio.Reader) map[int][]int {
	var n, m int
	fmt.Fscan(reader, &n, &m)

	graph := make(map[int][]int, n)
	for i := 0; i < m; i++ {
		var q int
		fmt.Fscan(reader, &q)

		slice := make([]int, q)
		for ind := 0; ind < q; ind++ {
			fmt.Fscan(reader, &slice[ind])
		}

		for len(slice) > 1 {
			start := slice[0]
			for index := 1; index < len(slice); index++ {
				graph[start] = append(graph[start], slice[index])
				graph[slice[index]] = append(graph[slice[index]], start)
			}
			slice = slice[1:]
		}
	}

	return graph
}
