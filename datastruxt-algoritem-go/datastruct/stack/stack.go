package stack

// Create by czx on 2019/12/25

type node struct {
	prev *node
	next *node
	data interface{}
}

func newNode(prev *node, next *node, data interface{}) *node {
	return &node{prev: prev, next: next, data: data}
}

//Stack 栈
type Stack struct {
	head *node
	tail *node
	size int
}

func NewStack() *Stack {
	tail := newNode(nil, nil, nil)
	head := newNode(nil, tail, nil)
	tail.prev = head
	return &Stack{head, tail, 0}
}

func (s *Stack) Push(e interface{}) {
	curr := s.head
	n := newNode(curr, curr.next, e)
	curr.next.prev = n
	curr.next = n
	s.size++
}

func (s *Stack) PopFront() (bool, interface{}) {
	if s.Empty() {
		return false, nil
	}
	curr := s.head.next
	curr.next.prev = curr.prev
	curr.prev.next = curr.next
	s.size--
	return true, curr.data
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Peek() interface{} {
	if s.Empty() {
		panic("queue is empty")
	}
	return s.head.next.data
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Clear() {
	s.head.next = s.tail
	s.tail.prev = s.head
	s.size = 0
}

func (s *Stack) Foreach(fn func(v interface{})) {
	curr := s.head.next
	for curr != s.tail {
		fn(curr.data)
		curr = curr.next
	}
}
