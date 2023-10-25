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
		case "front":
			stack.front()
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
	fmt.Println(stack.data[0])
	stack.data = stack.data[1:]
}

func (stack *stack) front() {
	if len(stack.data) == 0 {
		fmt.Println("error")
		return
	}
	fmt.Println(stack.data[0])
}

func (stack *stack) size() {
	fmt.Println(len(stack.data))
}

func (stack *stack) clear() {
	stack.data = []int{}
	fmt.Println("ok")
}
