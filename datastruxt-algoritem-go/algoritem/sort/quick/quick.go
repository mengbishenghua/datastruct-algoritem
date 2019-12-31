package quick

import (
	"com.datastruct.algoritem.go/algoritem/util"
)

// Create by czx on 2019/12/31

func Sort2Way(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l > r {
		return
	}
	pivot := partition(arr, l, r)
	quickSort(arr, l, pivot-1)
	quickSort(arr, pivot+1, r)

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

func Sort3way(arr []int) {
	sortThreeWay(arr, 0, len(arr)-1)
}

func sortThreeWay(arr []int, l, r int) {
	if l < r {
		lt, gt := sort3Way(arr, l, r)
		sortThreeWay(arr, l, lt-1)
		sortThreeWay(arr, gt, r)
	}
}

func sort3Way(arr []int, l, r int) (int, int) {
	//if r -l < 15 {
	//	insert.Sort(arr)
	//}
	/*pivot := arr[l]
	lt := l - 1
	gt := r + 1
	i := 1
	for i < gt {
		if arr[i] == pivot {
			i++
		} else if arr[i] < pivot {
			util.Swap(arr, i, lt+1)
			i++
			lt++
		} else {
			util.Swap(arr, i, gt-1)
			gt--
		}
	}
	sort3Way(arr, l, lt)
	sort3Way(arr, gt, r)*/
	v := arr[l]

	lt := l     // arr[l+1...lt] < v
	gt := r + 1 // arr[gt...r] > v
	i := l + 1  // arr[lt+1...i) == v 半开半闭
	for {
		switch {
		case i >= gt:
			// 这步未执行前,lt都是指向最后一个小于v的值,
			// 执行交换后,lt指向从左到右第一个等于v的值
			arr[l], arr[lt] = arr[lt], v
			return lt, gt
		case arr[i] == v:
			i++
		case arr[i] < v:
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		default:
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		}
	}

}
