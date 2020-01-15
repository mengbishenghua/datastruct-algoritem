package heap

import "com.datastruct.algoritem.go/algoritem/util"

// Create by czx on 2020/1/1

func Sort(arr []int) {
	heapSort(arr)
}

func heapSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	// 先转为大顶堆
	for i := 0; i < len(arr); i++ {
		Insert(arr, i)
	}

	// 将0号最大元素和尾部交换，之后从0号开始调整堆
	size := len(arr)
	util.Swap(arr, 0, len(arr)-1)
	size--
	for size > 0 {
		heapIfy(arr, 0, size)
		size--
		util.Swap(arr, 0, size)
	}
}

// 转为大顶堆
func Insert(arr []int, i int) {
	for arr[i] > arr[(i-1)/2] {
		util.Swap(arr, i, (i-1)/2)
		i = (i - 1) / 2
	}
}

// 从上往下修复堆的排序
func heapIfy(arr []int, index, size int) {
	left := 2*index + 1
	for left < size {
		max := left
		if left+1 < size && arr[left+1] > arr[left] {
			max = left + 1
		}
		if arr[max] < arr[index] {
			max = index
		}
		if max == index {
			break
		}
		util.Swap(arr, max, index)
		index = max
		left = 2*index + 1
	}
}
