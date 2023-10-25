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

	q_connections := make(map[int]int, n)
	graph := make(map[int][]int, n)
	graph_scan(q_connections, graph, m, n, reader)

	ans := topologicalSort(graph, q_connections, n)
	for i := 0; i < len(ans); i++ {
		if ans[i] == -1 {
			fmt.Print(ans[i])
		} else {
			fmt.Print(ans[i], " ")
		}
	}
}

func topologicalSort(graph map[int][]int, q_connections map[int]int, n int) []int {
	queue := make([]int, 0)

	for key, _ := range graph {
		if q_connections[key] == 0 {
			queue = append(queue, key)
		}
	}

	anslist := make([]int, 0)
	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]
		anslist = append(anslist, parent)

		for _, child := range graph[parent] {
			q_connections[child]--
			if q_connections[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	if len(anslist) != n {
		return []int{-1}
	} else {
		return anslist
	}
}

func graph_scan(q_connections map[int]int, graph map[int][]int, m, n int, reader *bufio.Reader) {
	for i := 0; i < m; i++ {
		var value1, value2 int
		fmt.Fscan(reader, &value1, &value2)
		graph[value1] = append(graph[value1], value2)
		q_connections[value2]++
	}
}
