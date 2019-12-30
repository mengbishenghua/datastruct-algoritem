package heap

import (
	"fmt"
	"testing"
)

// Create by czx on 2019/12/30

func Comparator(e1, e2 interface{}) int {
	n1, _ := e1.(int)
	n2, _ := e2.(int)
	return n1 - n2
}

func TestBinaryHeap(t *testing.T) {
	items := []interface{}{13, 14, 16, 19, 21, 19, 68, 65, 26, 32, 31}
	heap := NewBinaryHeap(Comparator, items)
	fmt.Println(heap.FindMin())
	heap.Insert(10)
	fmt.Println(heap.FindMin())
	fmt.Println(heap.DeleteMin())
}
