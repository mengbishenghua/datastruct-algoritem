package shell

import "com.datastruct.algoritem.go/algoritem/util"

// Create by czx on 2020/1/15

// 希尔排序：将数组分组，每组应用插入排序
func Sort(arr []int) {
	length := len(arr)
	if length == 1 {
		return
	}
	gap := length / 2

	for gap > 0 {
		for i := 0; i < gap; i++ {
			shellSort(arr, i, gap)
		}
		gap = gap / 2
	}
}

func shellSort(arr []int, start int, gap int) {
	// 插入排序
	for i := start + gap; i < len(arr); i += gap {
		temp := arr[i]
		j := i - gap
		for ; j >= 0 && arr[j] > temp; j -= gap {
			util.Swap(arr, j, j+gap)
		}
		temp, arr[j+gap] = arr[j+gap], temp
	}
}
