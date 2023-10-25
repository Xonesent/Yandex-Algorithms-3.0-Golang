package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(reader, &n, &k)

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i-j >= 0 {
				dp[i] += dp[i-j]
			}
		}
	}
	fmt.Println(dp[n])
}
