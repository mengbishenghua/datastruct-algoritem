package queue

import (
	"fmt"
	"testing"
)

// Create by czx on 2019/12/25

func TestQueue(t *testing.T) {
	dq := NewDequeue()
	for i := 0; i < 10; i++ {
		dq.Push(i)
	}
	dq.Foreach(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()
	fmt.Println("size: ", dq.Size())
	peek := dq.Peek()
	fmt.Println("peek value: ", peek)
	if ok, i := dq.Pop(); ok {
		fmt.Println("pop success: ", i)
	}
	if ok, i := dq.Pop(); ok {
		fmt.Println("pop success: ", i)
	}
	dq.Clear()
	fmt.Println("queue is empty: ", dq.Empty())

	for i := 0; i < 10; i++ {
		dq.PushFront(i)
	}
	dq.Foreach(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println("size: ", dq.Size())
	if ok, i := dq.PopFront(); ok {
		fmt.Println("pop front: ", i)
	}
	if ok, i := dq.Pop(); ok {
		fmt.Println("pop back: ", i)
	}
}
