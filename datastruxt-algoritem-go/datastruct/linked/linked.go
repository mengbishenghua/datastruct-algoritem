package linked

// Create by czx on 2019/12/25

type node struct {
	prev *node
	next *node
	data interface{}
}

func newNode(prev, next *node, data interface{}) *node {
	return &node{
		prev,
		next,
		data,
	}
}

type List struct {
	head *node
	tail *node
	size int
}

// TODO: 待实现
func NewList() *List {
	head := newNode(nil, nil, nil)
	tail := newNode(head, nil, nil)
	head.next = tail
	return &List{head: head, tail: tail}
}

func (l List) Append(e interface{}) {
	l.Insert(l.Size(), e)
}

func (l List) Insert(index int, e interface{}) {
	panic("implement me")
}

func (l List) Empty() bool {
	return l.Size() == 0
}

func (l List) Size() int {
	return l.size
}

func (l List) Get(index int) interface{} {
	panic("implement me")
}

func (l List) Set(index int, e interface{}) {
	panic("implement me")
}

func (l List) Remove(index int) (bool, interface{}) {
	panic("implement me")
}

func (l List) Clear() {
	panic("implement me")
}
