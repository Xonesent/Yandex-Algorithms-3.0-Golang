package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int = 0
	fmt.Fscan(reader, &n)

	arr := make([]int, n+1)
	ans := make_arr(n, arr)
	fmt.Println(ans - 1)
	anslist := find_ans_arr(n, ans, arr)
	for i := len(anslist) - 1; i >= 0; i-- {
		fmt.Print(anslist[i], " ")
	}
}

func make_arr(n int, arr []int) int {
	arr[1] = 1
	for i := 2; i < n+1; i++ {
		temp := arr[i-1] + 1

		if i%2 == 0 {
			if temp > arr[i/2]+1 {
				temp = arr[i/2] + 1
			}
		}
		if i%3 == 0 {
			if temp > arr[i/3]+1 {
				temp = arr[i/3] + 1
			}
		}

		arr[i] = temp
	}

	return arr[n]
}

func find_ans_arr(n int, ans int, arr []int) []int {
	anslist := make([]int, 0)
	for n != 0 {
		if n%3 == 0 && arr[n/3] == ans-1 {
			anslist = append(anslist, n)
			n /= 3
		} else if n%2 == 0 && arr[n/2] == ans-1 {
			anslist = append(anslist, n)
			n /= 2
		} else if arr[n-1] == ans-1 {
			anslist = append(anslist, n)
			n -= 1
		}
		ans--
	}

	return anslist
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
