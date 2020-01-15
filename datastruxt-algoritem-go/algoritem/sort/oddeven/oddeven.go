package oddeven

import "com.datastruct.algoritem.go/algoritem/util"

// Create by czx on 2020/1/15

// OddEvenSort奇偶排序
func Sort(arr []int) {
	// 是否已排好序,如果奇偶都没有发生交换，说明已排好序
	isSorted := false

	for isSorted == false {
		isSorted = true
		// 从偶数开始两两相邻比较
		for i := 0; i < len(arr)-1; i = i + 2 {
			if arr[i] > arr[i+1] {
				util.Swap(arr, i, i+1)
				isSorted = false
			}
		}

		// 从奇数开始两两相邻比较
		for j := 1; j < len(arr)-1; j = j + 2 {
			if arr[j] > arr[j+1] {
				util.Swap(arr, j, j+1)
				isSorted = false
			}
		}
	}

}
