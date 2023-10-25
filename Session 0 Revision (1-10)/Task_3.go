package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(reader, &n)

	stickers := make([]int, n)
	for i := 0; i < n; i++ {
		var value int
		fmt.Fscan(reader, &value)
		stickers[i] = value
	}

	sort.Ints(stickers)
	stickers = remove(stickers)

	fmt.Fscan(reader, &k)
	anslist := make([]int, k)
	for i := 0; i < k; i++ {
		var value int
		fmt.Fscan(reader, &value)
		anslist[i] = value
	}

	for _, value := range anslist {
		var left, right int = 0, len(stickers) - 1
		for left < right {
			mid := (left + right + 1) / 2
			if value <= stickers[mid] {
				right = mid - 1
			} else {
				left = mid
			}
		}

		if n > 0 {
			if value <= stickers[0] {
				fmt.Println(0)
			} else if value > stickers[len(stickers)-1] {
				fmt.Println(len(stickers))
			} else {
				fmt.Println(left + 1)
			}
		}
	}
}

func remove(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return nums[:j+1]
}
