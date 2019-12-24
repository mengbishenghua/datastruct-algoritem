package insert

import "com.datastruct.algoritem.go/algoritem/util"

// Create by czx on 2019/12/24

//Sort 插入排序
func Sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			util.Swap(arr, j, j+1)
		}
	}
}
