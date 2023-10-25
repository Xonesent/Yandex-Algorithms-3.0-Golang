package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	data []int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var operation string
	stack := stack{}

	for operation != "exit" {
		fmt.Fscan(reader, &operation)
		switch operation {
		case "push":
			var value int
			fmt.Fscan(reader, &value)
			stack.push(value)
		case "pop":
			stack.pop()
		case "back":
			stack.back()
		case "size":
			stack.size()
		case "clear":
			stack.clear()
		case "exit":
			fmt.Println("bye")
			break
		}
	}
}

func (stack *stack) push(n int) {
	stack.data = append(stack.data, n)
	fmt.Println("ok")
}

func (stack *stack) pop() {
	if len(stack.data) == 0 {
		fmt.Println("error")
		return
	}
	fmt.Println(stack.data[len(stack.data)-1])
	stack.data = stack.data[:len(stack.data)-1]
}

func (stack *stack) back() int {
	if len(stack.data) == 0 {
		fmt.Println("error")
		return -9999
	}
	fmt.Println(stack.data[len(stack.data)-1])
	return stack.data[len(stack.data)-1]
}

func (stack *stack) size() {
	fmt.Println(len(stack.data))
}

func (stack *stack) clear() {
	stack.data = []int{}
	fmt.Println("ok")
}
