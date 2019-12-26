package tree

// Create by czx on 2019/12/26

type node struct {
	data  interface{}
	left  *node
	right *node
}

func newNode(data interface{}, left *node, right *node) *node {
	return &node{data: data, left: left, right: right}
}

//BinarySearchTree 二叉搜索树
type BinarySearchTree struct {
	root *node
	size int
	fn   func(t1, t2 interface{}) int
}

func NewBinarySearchTree(fn func(t1, t2 interface{}) int) *BinarySearchTree {
	return &BinarySearchTree{fn: fn}
}

func (br *BinarySearchTree) Empty() bool {
	return br.root == nil
}

func (br *BinarySearchTree) Add(e interface{}) {
	br.root = br.add(br.root, e)
	br.size++
}

func (br *BinarySearchTree) FindMin() interface{} {
	return br.findMin(br.root).data
}

func (br *BinarySearchTree) FindMax() interface{} {
	return br.findMax(br.root).data
}

func (br *BinarySearchTree) Delete(e interface{}) {
	br.root = br.delete(br.root, e)
}

func (br *BinarySearchTree) InOrder(fn func(e interface{})) {
	br.inOrder(br.root, fn)
}

func (br *BinarySearchTree) inOrder(n *node, fn func(e interface{})) {
	if n == nil {
		return
	}
	if n.left != nil {
		br.inOrder(n.left, fn)
	}
	fn(n.data)
	if n.right != nil {
		br.inOrder(n.right, fn)
	}
}

func (br *BinarySearchTree) findMin(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (br *BinarySearchTree) findMax(n *node) *node {
	for n.right != nil {
		n = n.right
	}
	return n
}

func (br *BinarySearchTree) add(n *node, e interface{}) *node {
	if n == nil {
		return newNode(e, nil, nil)
	}
	comp := br.fn(e, n.data)
	if comp < 0 {
		n.left = br.add(n.left, e)
	} else if comp > 0 {
		n.right = br.add(n.right, e)
	}
	return n
}

func (br *BinarySearchTree) delete(n *node, e interface{}) *node {
	if n == nil {
		return n
	}
	comp := br.fn(e, n.data)
	if comp < 0 { // 小于当前节点
		n.left = br.delete(n.left, e)
	} else if comp > 0 { // 大于当前节点
		n.right = br.delete(n.right, e)
	} else {
		// 找到要删除的节点，如果有左右孩子
		if n.left != nil && n.right != nil {
			// 找到右子树的最小节点，实际上是找它的后继节点！
			min := br.findMin(n.right)
			// 删除最小节点，实际上是返回它的子节点，用父节点指向它的子节点
			n.right = br.delete(n.right, min.data)
			// 将后继节点的值赋值给当前节点
			n.data = min.data
		} else {
			// 只有一个孩子,返回它的孩子
			if n.left != nil {
				n = n.left
			} else {
				n = n.right
			}
		}
	}
	return n
}
