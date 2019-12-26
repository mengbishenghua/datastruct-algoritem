package tree

import (
	"fmt"
	"testing"
)

// Create by czx on 2019/12/26

func Comparator(t1, t2 interface{}) int {
	comp1, ok := t1.(int)
	if !ok {
		panic("convent fail")
	}
	comp2, ok := t2.(int)
	if !ok {
		panic("convent fail")
	}
	return comp1 - comp2
}

func TestBinarySearchTree(t *testing.T) {
	tr := NewBinarySearchTree(Comparator)
	tr.Add(6)
	tr.Add(2)
	tr.Add(8)
	tr.Add(1)
	tr.Add(5)
	tr.Add(3)
	tr.Add(4)

	tr.InOrder(func(e interface{}) {
		fmt.Print(e, " ")
	})
	fmt.Println("tr.Empty() = ", tr.Empty())

	max := tr.FindMax()
	fmt.Printf("max = %v\n", max)
	min := tr.FindMin()
	fmt.Printf("min = %v\n", min)

	tr.Delete(2)
	tr.InOrder(func(e interface{}) {
		fmt.Print(e, " ")
	})
	fmt.Println()
	tr.Delete(6)
	tr.InOrder(func(e interface{}) {
		fmt.Print(e, " ")
	})
}
