package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)
	stack := make([][2]int, 0)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		var value int
		fmt.Fscan(reader, &value)

		for len(stack) != 0 && value < stack[len(stack)-1][0] {
			ans[stack[len(stack)-1][1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, [2]int{value, i})
	}

	for i := 0; i < len(stack); i++ {
		ans[stack[i][1]] = -1
	}

	for i := 0; i < n; i++ {
		fmt.Print(ans[i], " ")
	}
}
