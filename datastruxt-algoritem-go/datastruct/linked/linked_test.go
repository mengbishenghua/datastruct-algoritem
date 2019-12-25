package linked

import (
	"fmt"
	"testing"
)

// Create by czx on 2019/12/25

func TestLinkedList(t *testing.T) {
	lk := NewList()
	for i := 0; i < 10; i++ {
		lk.Append(i)
	}
	lk.ForEach(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()
	fmt.Println("size: ", lk.Size())
	if lk.Size() != 10 {
		t.Error("size != ", 10)
	}
	fmt.Println("is empty: ", lk.Empty())
	lk.Insert(0, 100)
	lk.ForEach(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()
	fmt.Println(lk.Get(0))
	lk.Set(8, 4500)
	fmt.Println(lk.Get(8))
	if ok, i := lk.Remove(lk.Size() - 1); ok {
		fmt.Println("remove ok: ", i)
	}
	fmt.Println()
	lk.ForEach(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
}
