package main

import (
	"bufio"
	"fmt"
	"os"
)

type heap struct {
	data []int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	heap := &heap{data: []int{}}
	var operation string

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &operation)
		switch operation {
		case "0":
			var value int
			fmt.Fscan(reader, &value)
			heap.insert(value)
		case "1":
			ans, _ := heap.extract()
			fmt.Println(ans)
		}
	}
}

func (h *heap) insert(value int) {
	h.data = append(h.data, value)
	h.up(len(h.data) - 1)
}

func (h *heap) up(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[index] <= h.data[parent] {
			break
		}
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index = parent
	}
}

func (h *heap) extract() (int, error) {
	if len(h.data) == 0 {
		return 0, fmt.Errorf("error")
	}
	max := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down(0)
	return max, nil
}

func (h *heap) down(index int) {
	for {
		curr := index
		left := index*2 + 1
		right := index*2 + 2
		if left < len(h.data) && h.data[left] > h.data[curr] {
			curr = left
		}
		if right < len(h.data) && h.data[right] > h.data[curr] {
			curr = right
		}
		if curr == index {
			break
		}
		h.data[curr], h.data[index] = h.data[index], h.data[curr]
		index = curr
	}
}
