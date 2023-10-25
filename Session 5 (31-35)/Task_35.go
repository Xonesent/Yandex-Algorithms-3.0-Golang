package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	graph := make(map[int][]int)
	matrix_scan(graph)

	ans := find_ans(graph)

	if len(ans) == 0 {
		fmt.Print("NO")
	} else {
		ansmap := make(map[int]int)
		for i := 0; i < len(ans); i++ {
			_, ok := ansmap[ans[i]]
			if !ok {
				ansmap[ans[i]] = i
			} else {
				ans = ans[ansmap[ans[i]]+1:]
				break
			}
		}
		fmt.Println("YES")
		fmt.Println(len(ans))
		for i := 0; i < len(ans); i++ {
			fmt.Print(ans[i], " ")
		}
	}
}

func find_ans(graph map[int][]int) []int {
	visited := make(map[int]bool)
	for key := range graph {
		curr_visited := make(map[int]bool)
		curr := make([]int, 0)
		if !visited[key] {
			curr_visited[key] = true
			visited[key] = true
			curr = append(curr, key)
			stack := []int{key}

			for len(stack) > 0 {
				parent := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				for _, value := range graph[parent] {
					if !curr_visited[value] {
						visited[value] = true
						curr_visited[value] = true
						curr = append(curr, value)
						stack = append(stack, value)
					} else {
						if len(curr) > 2 && value != curr[len(curr)-2] && value != parent {
							curr = append(curr, value)
							return curr
						}
					}
				}
			}
		}
	}
	return []int{}
}

func matrix_scan(graph map[int][]int) {
	reader := bufio.NewReader(os.Stdin)

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
