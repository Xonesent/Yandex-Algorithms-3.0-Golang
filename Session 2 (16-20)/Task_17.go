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

	player1 := stack{data: []int{}}
	for i := 0; i < 5; i++ {
		var value int
		fmt.Fscan(reader, &value)
		player1.data = append(player1.data, value)
	}
	player2 := stack{data: []int{}}
	for i := 0; i < 5; i++ {
		var value int
		fmt.Fscan(reader, &value)
		player2.data = append(player2.data, value)
	}

	var ans int = 0
	for len(player1.data) != 0 && len(player2.data) != 0 && ans < 1000001 {
		card1 := player1.pop()
		card2 := player2.pop()
		if (card1 > card2 && !(card1 == 9 && card2 == 0)) || (card1 == 0 && card2 == 9) {
			player1.push(card1)
			player1.push(card2)
		} else {
			player2.push(card1)
			player2.push(card2)
		}

		ans++
	}

	if len(player1.data) == 0 {
		fmt.Println("second", ans)
	} else if len(player2.data) == 0 {
		fmt.Println("first", ans)
	} else {
		fmt.Println("botwa")
	}
}

func (stack *stack) push(n int) {
	stack.data = append(stack.data, n)
}

func (stack *stack) pop() int {
	var return_value int = stack.data[0]
	stack.data = stack.data[1:]
	return return_value
}
