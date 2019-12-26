package avltree

import (
	"fmt"
	"math"
)

// Create by czx on 2019/12/26

type node struct {
	value interface{}
	left  *node
	right *node
	h     int
}

func newNode(value interface{}, left *node, right *node, h int) *node {
	return &node{value: value, left: left, right: right, h: h}
}

type AvlTree struct {
	root *node
	fn   func(c1, c2 interface{}) int
}

func NewAvlTree(fn func(c1, c2 interface{}) int) *AvlTree {
	return &AvlTree{fn: fn}
}

func (at *AvlTree) Insert(value interface{}) {
	at.root = at.insert(at.root, value)
}

func (at *AvlTree) Remove(value interface{}) {
	at.root = at.remove(at.root, value)
}

func (at *AvlTree) InOrder() {
	at.inOrder(at.root)
	fmt.Println()
}

func (at *AvlTree) insert(n *node, value interface{}) *node {
	if n == nil {
		return newNode(value, nil, nil, 0)
	}
	comp := at.fn(value, n.value)
	if comp < 0 {
		n.left = at.insert(n.left, value)
	} else if comp > 0 {
		n.right = at.insert(n.right, value)
	}
	return balance(n)
}

func (at *AvlTree) remove(n *node, value interface{}) *node {
	if n == nil {
		return n
	}
	comp := at.fn(value, n.value)
	if comp < 0 {
		n.left = at.remove(n.left, value)
	} else if comp > 0 {
		n.right = at.remove(n.right, value)
	} else if n.left != nil && n.right != nil {
		min := at.findMin(n.right)
		n.right = at.remove(n.right, min)
		n.value = min
	} else {
		if n.left != nil {
			n = n.left
		} else {
			n = n.right
		}
	}
	return balance(n)
}

func (at *AvlTree) inOrder(n *node) {
	if n == nil {
		return
	}

	if n.left != nil {
		at.inOrder(n.left)
	}
	fmt.Print(n.value, " ")
	if n.right != nil {
		at.inOrder(n.right)
	}
}

func (at *AvlTree) findMin(n *node) interface{} {
	for n.left != nil {
		n = n.left
	}
	return n.value
}

// 平衡因子函数，调整平衡
func balance(n *node) *node {
	if n == nil {
		return nil
	}
	// 如果左子树大于右子树,说明插入的节点是在左子树
	if (maxH(n.left) - maxH(n.right)) > 1 {
		// 如果左子树的左子树高度大于左子树的右子树高度，单旋转
		if maxH(n.left.left) >= maxH(n.left.right) {
			n = rotateWithLeft(n)
		} else {
			n = doubleWithLeft(n)
		}
	} else if (maxH(n.right) - maxH(n.left)) > 1 { // 插入的节点在右子树
		if maxH(n.right.right) >= maxH(n.right.left) {
			n = rotateWithRight(n)
		} else {
			n = doubleWithRight(n)
		}
	}
	n.h = max(maxH(n.left), maxH(n.right)) + 1
	return n
}

// 左单旋
func rotateWithLeft(k2 *node) *node {
	k1 := k2.left
	k2.left = k1.right
	k1.right = k2
	k2.h = max(maxH(k2.left), maxH(k2.right)) + 1
	k1.h = max(maxH(k1.right), k2.h) + 1
	return k1
}

// 右单旋
func rotateWithRight(k1 *node) *node {
	k2 := k1.right
	k1.right = k2.left
	k2.left = k1
	k1.h = max(maxH(k1.left), maxH(k1.right)) + 1
	k2.h = max(maxH(k2.right), k1.h) + 1
	return k2
}

// 左双旋
func doubleWithLeft(k3 *node) *node {
	k3.left = rotateWithRight(k3.left)
	return rotateWithLeft(k3)
}

// 右双旋
func doubleWithRight(k2 *node) *node {
	k2.right = rotateWithLeft(k2.right)
	return rotateWithRight(k2)
}

func max(n1, n2 int) int {
	return int(math.Max(float64(n1), float64(n2)))
}

// 获取最大高度
func maxH(n *node) int {
	if n == nil {
		return -1
	} else {
		return n.h
	}
}
