package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var line string
	fmt.Fscan(reader, &line)

	var stack []rune
	var check bool = true
	for _, value := range line {
		if value == '(' || value == '[' || value == '{' {
			stack = append(stack, value)
		} else {
			switch value {
			case ']':
				if len(stack) == 0 || stack[len(stack)-1] != '[' {
					check = false
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			case '}':
				if len(stack) == 0 || stack[len(stack)-1] != '{' {
					check = false
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			case ')':
				if len(stack) == 0 || stack[len(stack)-1] != '(' {
					check = false
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
	}

	if len(stack) != 0 {
		check = false
	}

	if check {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
