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
		ans, ans_path := findShortcut(graph, start, end)
		fmt.Println(ans)
		for i := 0; i < len(ans_path); i++ {
			fmt.Print(ans_path[i], " ")
		}
	}
}

func reverse(anslist []int) {
	for i := 0; i < len(anslist)/2; i++ {
		anslist[i], anslist[len(anslist)-1-i] = anslist[len(anslist)-1-i], anslist[i]
	}
}

func find_path_way(pathway map[int]int, start, end int) []int {
	curr := end
	anslist := []int{curr}

	for pathway[curr] != 0 {
		anslist = append(anslist, pathway[curr])
		curr = pathway[curr]
	}

	reverse(anslist)
	return anslist
}

func findShortcut(graph map[int][]int, start, end int) (int, []int) {
	visited := make(map[int]bool)
	queue := make(map[int][]int)
	pathway := make(map[int]int)
	wave := 0

	visited[start] = true
	queue[wave] = append(queue[wave], start)
	pathway[start] = 0

	for wave <= len(graph) {
		for len(queue[wave]) > 0 {
			parent := queue[wave][0]
			queue[wave] = queue[wave][1:]

			for _, child := range graph[parent] {
				if !visited[child] {
					pathway[child] = parent
					if child == end {
						return wave + 1, find_path_way(pathway, start, end)
					}
					visited[child] = true
					queue[wave+1] = append(queue[wave+1], child)
				}
			}
		}
		wave++
	}

	return -1, []int{}
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
