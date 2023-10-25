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

	graph := make(map[int][]int, n)
	graph_scan(graph, m, n, reader)

	visited := make(map[int]bool, n)
	ansmatrix := make([][]int, 0)

	for key := range graph {
		if len(graph[key]) == 0 {
			ansmatrix = append(ansmatrix, []int{key})
		} else {
			candidate := poisk(graph, visited, key)
			if len(candidate) > 0 {
				ansmatrix = append(ansmatrix, candidate)
			}
		}
	}

	fmt.Println(len(ansmatrix))
	for _, value := range ansmatrix {
		fmt.Println(len(value))
		for i := 0; i < len(value); i++ {
			fmt.Print(value[i], " ")
		}
		fmt.Println()
	}
}

func graph_scan(graph map[int][]int, m, n int, reader *bufio.Reader) {
	scanned := make(map[int]bool, n)
	for i := 1; i <= n; i++ {
		scanned[i] = false
	}
	for i := 0; i < m; i++ {
		var value1, value2 int
		fmt.Fscan(reader, &value1, &value2)
		if value1 != value2 {
			graph[value1] = append(graph[value1], value2)
			graph[value2] = append(graph[value2], value1)
			scanned[value1] = true
			scanned[value2] = true
		}
	}
	for key, value := range scanned {
		if !value {
			graph[key] = []int{}
		}
	}
}

func poisk(graph map[int][]int, visited map[int]bool, search_value int) []int {
	if visited[search_value] != true {
		visited[search_value] = true
		anslist := []int{search_value}
		curr_poisk(graph, visited, &anslist, search_value)
		return anslist
	}
	return []int{}
}

func curr_poisk(graph map[int][]int, visited map[int]bool, curr_anslist *[]int, parent int) {
	for _, value := range graph[parent] {
		if visited[value] != true {
			visited[value] = true
			*curr_anslist = append(*curr_anslist, value)
			curr_poisk(graph, visited, curr_anslist, value)
		}
	}
}
