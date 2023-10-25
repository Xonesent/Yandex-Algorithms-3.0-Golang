package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int = 0
	fmt.Fscan(reader, &n)

	matrix := make([][3]int, n)
	for i := 0; i < n; i++ {
		var value1, value2, value3 int
		fmt.Fscan(reader, &value1, &value2, &value3)
		matrix[i] = [3]int{value1, value2, value3}
	}
	fmt.Println(find_ans(matrix, n))
}

func find_ans(matrix [][3]int, n int) int {
	dp := make([]int, n)
	if len(matrix) == 1 {
		return matrix[0][0]
	}
	dp[0] = matrix[0][0]
	if len(matrix) == 2 {
		return min([]int{matrix[0][0] + matrix[1][0], matrix[0][1]})
	}
	dp[1] = min([]int{matrix[0][0] + matrix[1][0], matrix[0][1]})
	if len(matrix) == 3 {
		return min([]int{matrix[0][2], dp[1] + matrix[2][0], dp[0] + matrix[1][1]})
	}
	dp[2] = min([]int{matrix[0][2], dp[1] + matrix[2][0], dp[0] + matrix[1][1]})
	for i := 3; i < n; i++ {
		dp[i] = min([]int{dp[i-3] + matrix[i-2][2], dp[i-2] + matrix[i-1][1], dp[i-1] + matrix[i][0]})
	}
	return dp[n-1]
}

func min(arr []int) int {
	var min int = math.MaxInt64
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
