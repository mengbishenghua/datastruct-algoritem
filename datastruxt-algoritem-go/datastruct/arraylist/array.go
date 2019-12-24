package arraylist

import (
	"errors"
	"math"
)

// Create by czx on 2019/12/25

const (
	DefaultCapacity = 16
)

type ArrayList struct {
	element  []interface{}
	size     int
	capacity int
}

func New() *ArrayList {
	return &ArrayList{
		element:  make([]interface{}, 0, DefaultCapacity),
		capacity: DefaultCapacity,
	}
}

func NewArrayList(capacity int) *ArrayList {
	if capacity <= 0 || capacity >= math.MaxInt64 {
		panic(errors.New("capacity out of bound"))
	}
	capacity = int(math.Max(float64(capacity), float64(DefaultCapacity)))
	return &ArrayList{
		element:  make([]interface{}, 0, capacity),
		capacity: capacity,
	}
}

func (l *ArrayList) Append(e interface{}) {
	l.Insert(l.size, e)
}

func (l *ArrayList) Insert(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("index out of bound")
	}
	if l.Size() == l.Capacity() {
		// 扩容
	}
	if index == l.size {
		l.element = append(l.element, e)
	} else {
		copy(l.element[index:], l.element[index+1:])
		l.element[index] = e
	}
	l.size++
}

func (l *ArrayList) Empty() bool {
	return l.Size() == 0
}

func (l *ArrayList) Size() int {
	return l.size
}

func (l *ArrayList) Capacity() int {
	return l.capacity
}

func (l *ArrayList) Get(index int) interface{} {
	l.checked(index)
	return l.element[index]
}

func (l *ArrayList) Set(index int, e interface{}) {
	l.checked(index)
	l.element[index] = e
}

func (l *ArrayList) Remove(index int) (bool, interface{}) {
	if index < 0 || index >= l.Size() {
		return false, nil
	}
	old := l.element[index]
	l.element = append(l.element[:index], l.element[index+1:]...)
	l.size--
	return true, old
}

func (l *ArrayList) Clear() {
	l.element = make([]interface{}, 0, DefaultCapacity)
	l.size = 0
	l.capacity = DefaultCapacity
}

func (l *ArrayList) ForEach(fn func(v interface{})) {
	for _, v := range l.element {
		fn(v)
	}
}

func (l *ArrayList) checked(index int) {
	if index < 0 || index >= l.Size() {
		panic("index out of bound")
	}
}
