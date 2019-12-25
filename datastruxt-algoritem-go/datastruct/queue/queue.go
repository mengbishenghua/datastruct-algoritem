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

//Queue  双向链表实现的队列
type Queue struct {
	head *node
	tail *node
	size int
}

func NewQueue() *Queue {
	tail := newNode(nil, nil, nil)
	head := newNode(nil, tail, nil)
	tail.prev = head
	return &Queue{head, tail, 0}
}

func (q *Queue) Push(e interface{}) {
	curr := q.tail
	n := newNode(curr.prev, curr, e)
	curr.prev.next = n
	curr.prev = n
	q.size++
}

func (q *Queue) Pop() (bool, interface{}) {
	if q.Empty() {
		return false, nil
	}
	curr := q.head.next
	curr.next.prev = curr.prev
	curr.prev.next = curr.next
	q.size--
	return true, curr.data
}

func (q *Queue) Empty() bool {
	return q.Size() == 0 || q.head == q.tail
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Peek() interface{} {
	if q.Empty() {
		panic("queue is empty")
	}
	return q.head.next.data
}

func (q *Queue) Clear() {
	q.head.next = q.tail
	q.tail.prev = q.head
	q.size = 0
}

func (q *Queue) Foreach(fn func(v interface{})) {
	curr := q.head.next
	for curr != q.tail {
		fn(curr.data)
		curr = curr.next
	}
}
