package array

import (
	"errors"
	"math"
)

// Create by czx on 2019/12/25

const (
	DefaultCapacity = 16
)

type Array struct {
	element  []interface{}
	size     int
	capacity int
}

func New() *Array {
	return &Array{
		element:  make([]interface{}, 0, DefaultCapacity),
		capacity: DefaultCapacity,
	}
}

func NewArrayList(capacity int) *Array {
	if capacity <= 0 || capacity >= math.MaxInt64 {
		panic(errors.New("capacity out of bound"))
	}
	capacity = int(math.Max(float64(capacity), float64(DefaultCapacity)))
	return &Array{
		element:  make([]interface{}, 0, capacity),
		capacity: capacity,
	}
}

func (l *Array) Append(e interface{}) {
	l.Insert(l.size, e)
}

func (l *Array) Insert(index int, e interface{}) {
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

func (l *Array) Empty() bool {
	return l.Size() == 0
}

func (l *Array) Size() int {
	return l.size
}

func (l *Array) Capacity() int {
	return l.capacity
}

func (l *Array) Get(index int) interface{} {
	l.checked(index)
	return l.element[index]
}

func (l *Array) Set(index int, e interface{}) {
	l.checked(index)
	l.element[index] = e
}

func (l *Array) Remove(index int) (bool, interface{}) {
	if index < 0 || index >= l.Size() {
		return false, nil
	}
	old := l.element[index]
	l.element = append(l.element[:index], l.element[index+1:]...)
	l.size--
	return true, old
}

func (l *Array) Clear() {
	l.element = make([]interface{}, 0, DefaultCapacity)
	l.size = 0
	l.capacity = DefaultCapacity
}

func (l *Array) ForEach(fn func(v interface{})) {
	for _, v := range l.element {
		fn(v)
	}
}

func (l *Array) checked(index int) {
	if index < 0 || index >= l.Size() {
		panic("index out of bound")
	}
}
