package quick

import "com.datastruct.algoritem.go/algoritem/util"

// Create by czx on 2019/12/31

func Sort2Way(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l < r {
		pivot := partition(arr, l, r)
		quickSort(arr, l, pivot-1)
		quickSort(arr, pivot+1, r)
	}

}

func partition(arr []int, l int, r int) int {
	pivot := arr[l]
	left := l + 1
	right := r

	for {
		// left == right 也需要进行判断
		for left <= right && arr[left] < pivot {
			left++
		}
		for left <= right && arr[right] > pivot {
			right--
		}
		if left > right {
			break
		} else {
			util.Swap(arr, left, right)
			left++
			right--
		}
	}
	util.Swap(arr, l, right)
	return right

	//for {
	//	switch {
	//	case left > right:
	//		util.Swap(arr, l, right)
	//		return right
	//	case arr[left] < pivot:
	//		left++
	//	case arr[right] > pivot:
	//		right--
	//	default:
	//		util.Swap(arr, left, right)
	//		left++
	//		right--
	//	}
	//}
}
