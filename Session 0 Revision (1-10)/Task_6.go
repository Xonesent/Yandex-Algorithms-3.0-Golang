package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var m, n, ans int
	fmt.Fscan(reader, &m, &n)
	systems := make([][2]int, 0)

	for i := 0; i < n; i++ {
		var left, right int
		fmt.Fscan(reader, &left, &right)
		if left > m || left < 0 || right > m || right < 0 {
			continue
		} else {
			systems = append(systems, [2]int{left, right})
		}
	}

	for i := 0; i < len(systems)/2; i++ {
		systems[i], systems[len(systems)-1-i] = systems[len(systems)-1-i], systems[i]
	}

	anslist := make([][2]int, 0)
	for i := 0; i < len(systems); i++ {
		var check bool = true
		for ind := 0; ind < len(anslist) && check; ind++ {
			if (anslist[ind][0] <= systems[i][0] && systems[i][0] <= anslist[ind][1]) || (anslist[ind][0] <= systems[i][1] && systems[i][1] <= anslist[ind][1]) || (systems[i][1] >= anslist[ind][1] && systems[i][0] <= anslist[ind][0]) {
				check = false
			}
		}
		if check {
			ans++
		}
		anslist = append(anslist, systems[i])
	}

	fmt.Println(ans)
}
