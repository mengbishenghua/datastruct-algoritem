package bubble

import "com.datastruct.algoritem.go/algoritem/util"

// Create by czx on 2019/12/24

//Sort 冒泡排序
func Sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				util.Swap(arr, j, j+1)
			}
		}
	}
}
