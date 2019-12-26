package avltree

import "testing"

// Create by czx on 2019/12/26

func Compare(c1, c2 interface{}) int {
	v1, _ := c1.(int)
	v2, _ := c2.(int)
	return v1 - v2
}

func TestAvlTree(t *testing.T) {
	at := NewAvlTree(Compare)
	at.Insert(6)
	at.Insert(2)
	at.Insert(8)
	at.Insert(1)
	at.Insert(5)
	at.Insert(3)
	at.Insert(4)

	at.InOrder()

	at.Remove(2)
	at.Remove(5)
	at.InOrder()
}
