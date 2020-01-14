package queue

import "math"

// Create by czx on 2019/12/25
const DefaultCapacity = 16

// RingQueue 循环队列
// 添加一个元素时，rear++， 删除一个元素时，front++
// 队列为空时: front == rear
// 队列为满： (rear + 1) % cap(element) == front
// 队列大小：（rear-front + cap(element)) % cap(element)
// 下一个索引位置：(front/rear+1) % cap(element)
type RingQueue struct {
	element []interface{}
	front   int
	rear    int
}

func NewRingQueue(n int) *RingQueue {
	if n <= 0 || n > math.MaxInt64 {
		panic("n overflow")
	}
	return &RingQueue{element: make([]interface{}, n, n)}
}

func (r *RingQueue) Push(e interface{}) {
	if r.full() {
		r.resize(cap(r.element) * 2)
	}
	r.element[r.rear] = e
	r.rear = r.nextIndex(r.rear)
}

func (r *RingQueue) Pop() (bool, interface{}) {
	if r.Empty() {
		return false, nil
	}
	res := r.element[r.front]
	r.front = r.nextIndex(r.front)
	return true, res
}

func (r *RingQueue) Empty() bool {
	return r.front == r.rear
}

func (r *RingQueue) full() bool {
	return (r.rear+1)%cap(r.element) == r.front
}

func (r *RingQueue) Size() int {
	return (r.rear - r.front + cap(r.element)) % cap(r.element)
}

func (r *RingQueue) Clear() {
	r.front = 0
	r.rear = 0
}

func (r *RingQueue) resize(capacity int) {
	newElement := make([]interface{}, capacity)
	curr := r.front
	for i := 0; curr != r.rear; i++ {
		newElement[i] = r.element[curr]
		curr = r.nextIndex(curr)
	}
	r.front = 0
	r.rear = curr
	r.element = newElement
}

func (r *RingQueue) Foreach(fn func(v interface{})) {
	curr := r.front
	for curr != r.rear {
		fn(r.element[curr])
		curr = r.nextIndex(curr)
	}
}

// 计算下一个索引的位置
func (r *RingQueue) nextIndex(curr int) int {
	return (curr + 1) % cap(r.element)
}
