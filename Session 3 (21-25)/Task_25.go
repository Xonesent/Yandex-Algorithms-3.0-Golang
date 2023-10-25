package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int = 0
	fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		var value int
		fmt.Fscan(reader, &value)
		arr[i] = value
	}
	sort.Ints(arr)

	fmt.Println(find_ans(arr, n))
}

func find_ans(arr []int, n int) int {
	if len(arr) == 2 {
		return arr[1] - arr[0]
	}
	if len(arr) == 3 {
		return arr[2] - arr[0]
	}
	if len(arr) == 4 {
		return arr[1] - arr[0] + arr[3] - arr[2]
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = arr[1] - arr[0]
	dp[2] = arr[2] - arr[0]
	dp[3] = arr[1] - arr[0] + arr[3] - arr[2]
	for i := 4; i < n; i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + arr[i] - arr[i-1]
	}

	return dp[n-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
