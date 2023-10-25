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

	if find_ans(graph) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func graph_scan(graph map[int][]int, m, n int, reader *bufio.Reader) {
	for i := 0; i < m; i++ {
		var value1, value2 int
		fmt.Fscan(reader, &value1, &value2)
		graph[value1] = append(graph[value1], value2)
		graph[value2] = append(graph[value2], value1)
	}
}

func find_ans(graph map[int][]int) bool {
	status := make(map[int]int)

	for key := range graph {
		if status[key] == 0 {
			status[key] = 1
			stack := []int{key}

			for len(stack) > 0 {
				parent := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				for _, child := range graph[parent] {
					if status[child] == 0 {
						status[child] = -1 * status[parent]
						stack = append(stack, child)
					} else if status[child] == status[parent] {
						return false
					}
				}
			}
		}
	}

	return true
}
