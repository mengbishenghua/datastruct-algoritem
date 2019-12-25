package queue

// Create by czx on 2019/12/25

type node struct {
	prev *node
	next *node
	data interface{}
}

func newNode(prev *node, next *node, data interface{}) *node {
	return &node{prev: prev, next: next, data: data}
}

//DeQueue  双向链表实现的队列
type DeQueue struct {
	head *node
	tail *node
	size int
}

func NewDequeue() *DeQueue {
	tail := newNode(nil, nil, nil)
	head := newNode(nil, tail, nil)
	tail.prev = head
	return &DeQueue{head, tail, 0}
}

func (q *DeQueue) Push(e interface{}) {
	curr := q.tail
	n := newNode(curr.prev, curr, e)
	curr.prev.next = n
	curr.prev = n
	q.size++
}

func (q *DeQueue) PushFront(e interface{}) {
	curr := q.head
	n := newNode(curr, curr.next, e)
	curr.next.prev = n
	curr.next = n
	q.size++
}

func (q *DeQueue) Pop() (bool, interface{}) {
	if q.Empty() {
		return false, nil
	}
	curr := q.tail.prev
	curr.prev.next = curr.next
	curr.next.prev = curr.prev
	q.size--
	return true, curr.data
}

func (q *DeQueue) PopFront() (bool, interface{}) {
	if q.Empty() {
		return false, nil
	}
	curr := q.head.next
	curr.next.prev = curr.prev
	curr.prev.next = curr.next
	q.size--
	return true, curr.data
}

func (q *DeQueue) Empty() bool {
	return q.Size() == 0 || q.head == q.tail
}

func (q *DeQueue) Size() int {
	return q.size
}

func (q *DeQueue) Peek() interface{} {
	if q.Empty() {
		panic("queue is empty")
	}
	return q.head.next.data
}

func (q *DeQueue) Clear() {
	q.head.next = q.tail
	q.tail.prev = q.head
	q.size = 0
}

func (q *DeQueue) Foreach(fn func(v interface{})) {
	curr := q.head.next
	for curr != q.tail {
		fn(curr.data)
		curr = curr.next
	}
}
