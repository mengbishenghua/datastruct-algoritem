package queue

import "math"

// Create by czx on 2019/12/25
const DefaultCapacity = 16

// RingQueue 循环队列, 我们多创建一个容量便于判断是否队满：make(type, len, cap+1)
// 添加一个元素时，rear++， 删除一个元素时，front++
// 队列为空时: front == rear
// 队列为满： (rear + 1) % cap(element) == front
// 队列大小：Math.abs(rear-front) % cap(element)
// 下一个索引位置：(front/rear+1) % cap(element)
type RingQueue struct {
	element []interface{}
	front   int
	rear    int
	size    int
}

func NewRingQueue(n int) *RingQueue {
	if n <= 0 || n > math.MaxInt64 {
		panic("n overflow")
	}
	return &RingQueue{element: make([]interface{}, n, n+1)}
}

func (r *RingQueue) Push(e interface{}) {
	if r.full() {
		r.resize(cap(r.element)*2 - 1)
	}
	r.element[r.rear] = e
	r.rear = r.nextIndex(r.rear)
	r.size++
}

func (r *RingQueue) Pop() (bool, interface{}) {
	if r.Empty() {
		return false, nil
	}
	res := r.element[r.front]
	r.front = r.nextIndex(r.front)
	r.size--
	return true, res
}

func (r *RingQueue) Empty() bool {
	return r.front == r.rear
}

func (r *RingQueue) full() bool {
	return (r.rear+1)%cap(r.element) == r.front
}

func (r *RingQueue) Size() int {
	// size = int(math.Abs(float64(r.rear)-float64(r.front))) % cap(r.element)
	return r.size
}

func (r *RingQueue) Clear() {
	r.element = make([]interface{}, DefaultCapacity, DefaultCapacity+1)
	r.size = 0
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
	r.rear = r.Size()
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
