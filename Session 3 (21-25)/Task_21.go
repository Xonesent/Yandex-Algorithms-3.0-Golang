package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	dp := make([]int, 35)
	dp[0] = 2
	dp[1] = 4
	dp[2] = 7
	for i := 3; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	fmt.Print(dp[n-1])
}
