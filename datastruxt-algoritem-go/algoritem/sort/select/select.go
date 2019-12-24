package sort

import "com.datastruct.algoritem.go/algoritem/util"

// Create by czx on 2019/12/24

func Sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = j
			}
		}
		if min != i {
			util.Swap(arr, min, i)
		}
	}
}
