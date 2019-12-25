package queue

import (
	"fmt"
	"testing"
)

// Create by czx on 2019/12/25

func TestQueue(t *testing.T) {
	q := NewQueue()
	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	q.Foreach(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()
	fmt.Println("size: ", q.Size())
	peek := q.Peek()
	fmt.Println("peek value: ", peek)
	if ok, i := q.Pop(); ok {
		fmt.Println("pop success: ", i)
	}
	if ok, i := q.Pop(); ok {
		fmt.Println("pop success: ", i)
	}
	q.Clear()
	fmt.Println("queue is empty: ", q.Empty())
}
