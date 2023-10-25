package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, ans int
	fmt.Fscan(reader, &n)

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &slice[i])
		if i != 0 {
			if slice[i] >= slice[i-1] {
				ans += slice[i-1]
			} else if slice[i] < slice[i-1] {
				ans += slice[i]
			}
		}
	}

	fmt.Println(ans)
}
