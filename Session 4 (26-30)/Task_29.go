package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	schedule := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &schedule[i])
	}
	matrix := make([][]int, n+1)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for ind := 0; ind < n; ind++ {
			matrix[i][ind] = math.MaxInt32
		}
	}

	if n == 0 {
		fmt.Println(0)
		fmt.Println(0, 0)
	} else if n == 1 {
		fmt.Println(schedule[0])
		if schedule[0] > 100 {
			fmt.Println(1, 0)
		} else {
			fmt.Println(0, 0)
		}
	} else {
		find_answers(n, matrix, schedule)
	}

	// for i := 0; i < n; i++ {
	// 	for ind := 0; ind < n; ind++ {
	// 		if matrix[i][ind] == 2147483647 {
	// 			fmt.Print(9999, " ")
	// 		} else {
	// 			fmt.Print(matrix[i][ind], " ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// for i := 0; i < n; i++ {
	// 	fmt.Print(schedule[i], " ")
	// }
}

func find_answers(n int, matrix [][]int, schedule []int) {
	matrix_filling(matrix, schedule, n)

	coupon1 := 0
	var ans int = math.MaxInt32
	for i := n - 1; i >= 0; i-- {
		if matrix[n-1][i] <= ans {
			ans = matrix[n-1][i]
			coupon1 = i
		}
	}

	anslist, coupon2 := find_coupon_day(matrix, schedule, n, coupon1)

	fmt.Println(ans)
	fmt.Println(coupon1, coupon2)
	for i := len(anslist) - 1; i >= 0; i-- {
		fmt.Print(anslist[i], " ")
	}
}

func find_coupon_day(matrix [][]int, schedule []int, n int, index int) ([]int, int) {
	var coupon2 int

	anslist := make([]int, 0)
	for i := n - 1; i > 0; i-- {
		if index+1 < n && matrix[i][index] == matrix[i-1][index+1] {
			coupon2++
			anslist = append(anslist, i+1)
			index++
		} else if index-1 >= 0 && matrix[i][index] == matrix[i-1][index-1]+schedule[i] {
			index--
		}
	}

	return anslist, coupon2
}

func matrix_filling(matrix [][]int, schedule []int, n int) {
	if schedule[0] > 100 {
		matrix[0][1] = schedule[0]
	} else {
		matrix[0][0] = schedule[0]
	}
	for i := 1; i < n; i++ {
		for ind := 0; ind < n; ind++ {
			if schedule[i] > 100 {
				if ind == 0 {
					matrix[i][ind] = matrix[i-1][ind+1]
				} else if ind < n-1 {
					matrix[i][ind] = min(matrix[i-1][ind-1]+schedule[i], matrix[i-1][ind+1])
				} else {
					matrix[i][ind] = matrix[i-1][ind-1] + schedule[i]
				}
			} else {
				if ind < n-1 {
					matrix[i][ind] = min(matrix[i-1][ind]+schedule[i], matrix[i-1][ind+1])
				} else {
					matrix[i][ind] = matrix[i-1][ind] + schedule[i]
				}
			}
		}
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
