package queue

import (
	"fmt"
	"testing"
)

// Create by czx on 2019/12/25

func TestRingQueue(t *testing.T) {
	r := NewRingQueue(10)
	for i := 0; i < 10; i++ {
		r.Push(i)
	}
	r.Foreach(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()
	fmt.Println("size: ", r.Size())
	r.Push(100)
	r.Foreach(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()
	fmt.Println("size: ", r.Size())
	fmt.Println("is empty: ", r.Empty())

	if ok, i := r.Pop(); ok {
		fmt.Println("pop success: ", i)
	}
	if ok, i := r.Pop(); ok {
		fmt.Println("pop success: ", i)
	}
	if ok, i := r.Pop(); ok {
		fmt.Println("pop success: ", i)
	}

	fmt.Println("size: ", r.Size())

	r.Clear()
	fmt.Println("is empty: ", r.Empty())
	for i := 0; i < 20; i++ {
		r.Push(i)
	}
	r.Foreach(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()
}
