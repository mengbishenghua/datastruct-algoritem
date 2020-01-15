package oddeven

import "com.datastruct.algoritem.go/algoritem/util"

// Create by czx on 2020/1/15

// OddEvenSort奇偶排序
func Sort(arr []int) {
	isSorted := false

	for isSorted == false {
		isSorted = true
		for i := 0; i < len(arr)-1; i = i + 1 {
			if arr[i] > arr[i+1] {
				util.Swap(arr, i, i+1)
				isSorted = false
			}
		}

		for j := 1; j < len(arr)-1; j = j + 2 {
			if arr[j] > arr[j+1] {
				util.Swap(arr, j, j+1)
				isSorted = false
			}
		}
	}

}
