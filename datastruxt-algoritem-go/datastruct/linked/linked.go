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

func NewList() *List {
	head := newNode(nil, nil, nil)
	tail := newNode(head, nil, nil)
	head.next = tail
	return &List{head: head, tail: tail}
}

func (l *List) Append(e interface{}) {
	l.Insert(l.Size(), e)
}

func (l *List) Insert(index int, e interface{}) {
	l.check(index)
	curr := l.move(index)
	n := &node{
		prev: curr,
		next: curr.next,
		data: e,
	}
	curr.next.prev = n
	curr.next = n
	l.size++
}

func (l *List) Empty() bool {
	return l.Size() == 0 || l.head == l.tail
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Get(index int) interface{} {
	if index < 0 || index >= l.Size() {
		panic("index out of bound")
	}
	curr := l.move(index)
	return curr.next.data
}

func (l *List) Set(index int, e interface{}) {
	if index < 0 || index >= l.Size() {
		panic("index out of bound")
	}
	curr := l.move(index)
	curr.next.data = e
}

func (l *List) Remove(index int) (bool, interface{}) {
	if index < 0 || index >= l.Size() {
		panic("index out of bound")
	}
	if l.Empty() {
		return false, nil
	}
	curr := l.move(index)
	old := curr.next
	old.next.prev = curr
	curr.next = old.next
	l.size--
	return true, old.data
}

func (l *List) Clear() {
	l.head.next = l.tail
	l.tail.prev = l.head
	l.size = 0
}

func (l *List) move(index int) (curr *node) {
	curr = l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return
}

func (l *List) check(index int) {
	if index < 0 || index > l.Size() {
		panic("index out of bound")
	}
}

func (l *List) ForEach(fn func(v interface{})) {
	curr := l.head.next
	for curr != l.tail {
		fn(curr.data)
		curr = curr.next
	}
}
