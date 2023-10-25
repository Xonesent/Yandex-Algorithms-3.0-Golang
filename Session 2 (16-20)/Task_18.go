package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

type Deque struct {
	Front  *Node
	Rear   *Node
	Length int
}

func (d *Deque) IsEmpty() bool {
	return d.Front == nil
}

func (d *Deque) AddFront(item int) {
	node := &Node{Value: item, Prev: nil, Next: nil}
	if d.IsEmpty() {
		d.Front = node
		d.Rear = node
		d.Length = 1
	} else {
		d.Length += 1
		node.Prev = d.Front
		d.Front.Next = node
		d.Front = node
	}
}

func (d *Deque) AddRear(item int) {
	node := &Node{Value: item, Prev: nil, Next: nil}
	if d.IsEmpty() {
		d.Front = node
		d.Rear = node
		d.Length = 1
	} else {
		d.Length += 1
		node.Next = d.Rear
		d.Rear.Prev = node
		d.Rear = node
	}
}

func (d *Deque) RemoveFront() (int, error) {
	if d.IsEmpty() {
		return 0, fmt.Errorf("error")
	}
	return_value := d.Front.Value
	d.Front = d.Front.Prev
	if d.Front == nil {
		d.Rear = nil
	} else {
		d.Front.Next = nil
	}
	d.Length -= 1
	return return_value, nil
}

func (d *Deque) RemoveRear() (int, error) {
	if d.IsEmpty() {
		return 0, fmt.Errorf("error")
	}
	return_value := d.Rear.Value
	d.Rear = d.Rear.Next
	if d.Rear == nil {
		d.Front = nil
	} else {
		d.Rear.Prev = nil
	}
	d.Length -= 1
	return return_value, nil
}

func (d *Deque) FrontValue() (int, error) {
	if d.IsEmpty() {
		return 0, fmt.Errorf("error")
	}
	return d.Front.Value, nil
}

func (d *Deque) RearValue() (int, error) {
	if d.IsEmpty() {
		return 0, fmt.Errorf("error")
	}
	return d.Rear.Value, nil
}

func (d *Deque) Clear() {
	for !d.IsEmpty() {
		d.Rear = d.Rear.Next
		if d.Rear == nil {
			d.Front = nil
		} else {
			d.Rear.Prev = nil
		}
		d.Length -= 1
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var operation string
	deque := &Deque{Front: nil, Rear: nil, Length: 0}

	for operation != "exit" {
		fmt.Fscan(reader, &operation)
		switch operation {
		case "push_front":
			var value int
			fmt.Fscan(reader, &value)
			deque.AddFront(value)
			fmt.Println("ok")
		case "push_back":
			var value int
			fmt.Fscan(reader, &value)
			deque.AddRear(value)
			fmt.Println("ok")
		case "pop_front":
			item, err := deque.RemoveFront()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(item)
			}
		case "pop_back":
			item, err := deque.RemoveRear()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(item)
			}
		case "front":
			item, err := deque.FrontValue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(item)
			}
		case "back":
			item, err := deque.RearValue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(item)
			}
		case "size":
			fmt.Println(deque.Length)
		case "clear":
			deque.Clear()
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			break
		}
	}
}
