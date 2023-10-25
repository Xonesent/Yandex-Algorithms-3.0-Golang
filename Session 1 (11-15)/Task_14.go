package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, count int
	fmt.Fscan(reader, &n)
	var check bool = true
	var curr int = 1
	dq := list.New()

	for count != n {
		var value int
		fmt.Fscan(reader, &value)

		if value == curr && check {
			curr++
			for dq.Len() != 0 && dq.Back().Value == curr {
				curr++
				dq.Remove(dq.Back())
			}
		} else {
			if dq.Len() != 0 && value > dq.Back().Value.(int) {
				check = false
			}
			dq.PushBack(value)
		}
		count++
	}

	if check && dq.Len() == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
