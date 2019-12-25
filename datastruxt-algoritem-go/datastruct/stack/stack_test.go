package stack

import (
	"fmt"
	"testing"
)

// Create by czx on 2019/12/25

func TestStack(t *testing.T) {
	s := NewStack()
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	s.Foreach(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println("size: ", s.Size())

	fmt.Println("peek: ", s.Peek())
	fmt.Println("s.Empty() = ", s.Empty())

	if ok, i := s.PopFront(); ok {
		fmt.Println("pop success: ", i)
	}
	if ok, i := s.PopFront(); ok {
		fmt.Println("pop success: ", i)
	}

	s.Clear()
	fmt.Println(" s.Size() = ", s.Size())
	fmt.Println("s.Empty() = ", s.Empty())
}
