package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var stack []int
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		switch {
		case char >= '0' && char <= '9':
			stack = append(stack, int(char-'0'))
		case char == '+':
			value1 := stack[len(stack)-2]
			value2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, value1+value2)
		case char == '-':
			value1 := stack[len(stack)-2]
			value2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, value1-value2)
		case char == '*':
			value1 := stack[len(stack)-2]
			value2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, value1*value2)
		}
	}

	fmt.Println(stack[len(stack)-1])
}
