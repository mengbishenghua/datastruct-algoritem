package heap

import (
	"errors"
)

// Create by czx on 2019/12/30

//优先队列：当前节点的左子树：2i, 右子树：2i+1
//优先队列父节点：i/2
//这里实现了最小堆
//这里的二叉堆根节点从1开始的，0号位置保留
type BinaryHeap struct {
	currSize int
	array    []interface{}
	fn       Compare
}

type Compare func(e1, e2 interface{}) int

func NewBinaryHeap(fn func(e1, e2 interface{}) int, array []interface{}) *BinaryHeap {
	heap := new(BinaryHeap)
	heap.fn = fn
	heap.currSize = len(array)
	heap.array = make([]interface{}, (heap.currSize+2)*11/10)
	i := 1
	for _, item := range array {
		heap.array[i] = item
		i++
	}
	heap.buildHeap()
	return heap
}

func (bh *BinaryHeap) Insert(e interface{}) {
	if bh.currSize == len(bh.array)-1 {
		bh.reSize(2*len(bh.array) - 1)
	}
	hole := bh.currSize + 1
	bh.currSize += 1
	for bh.array[0] = e; bh.fn(e, bh.array[hole/2]) < 0; hole /= 2 {
		bh.array[hole] = bh.array[hole/2]
	}
	bh.array[hole] = e
}

func (bh *BinaryHeap) Empty() bool {
	return bh.currSize == 0
}

func (bh *BinaryHeap) FindMin() interface{} {
	return bh.array[1]
}

func (bh *BinaryHeap) DeleteMin() (interface{}, error) {
	if bh.Empty() {
		return nil, errors.New("heap is empty")
	}
	min := bh.FindMin()
	bh.array[1] = bh.array[bh.currSize-1]
	bh.currSize--
	bh.shiftDown(1)

	return min, nil
}

func (bh *BinaryHeap) shiftDown(hole int) {
	child := 0
	tmp := bh.array[hole]
	for ; hole*2 <= bh.currSize; hole = child {
		child = 2 * hole
		if child != bh.currSize && bh.fn(bh.array[child+1], bh.array[child]) < 0 {
			child++
		}
		if bh.fn(bh.array[child], tmp) < 0 {
			bh.array[hole] = bh.array[child]
		} else {
			break
		}
	}
	bh.array[hole] = tmp
}

func (bh *BinaryHeap) buildHeap() {
	for i := bh.currSize / 2; i > 0; i-- {
		bh.shiftDown(i)
	}
}

func (bh *BinaryHeap) reSize(newSize int) {

}
