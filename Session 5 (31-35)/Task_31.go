package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(reader, &n, &m)

	graph := make(map[int][]int, n)
	for i := 0; i < m; i++ {
		var value1, value2 int
		fmt.Fscan(reader, &value1, &value2)
		if value1 != value2 {
			graph[value1] = append(graph[value1], value2)
			graph[value2] = append(graph[value2], value1)
		}
	}

	visited := make(map[int]bool)
	dfs(graph, visited, 1)

	ans := []int{}
	for key, value := range visited {
		if value {
			ans = append(ans, key)
		}
	}

	sort.Ints(ans)

	fmt.Println(len(ans))
	for i := 0; i < len(ans); i++ {
		fmt.Print(ans[i], " ")
	}
}

func dfs(graph map[int][]int, visited map[int]bool, search_value int) {
	visited[search_value] = true
	for _, value := range graph[search_value] {
		if !visited[value] {
			dfs(graph, visited, value)
		}
	}
}
