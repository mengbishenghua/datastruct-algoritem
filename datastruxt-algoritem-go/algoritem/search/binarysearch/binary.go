package binarysearch

import "fmt"

// Create by czx on 2019/12/24

//Search 二分搜索
func Search(arr []int, begin, end, x int) (bool, int) {
	n := 1
	if len(arr) == 0 {
		return false, -1
	}
	var mid = 0
	for begin <= end {
		n++
		fmt.Println("查找了", n, "次")
		// 中值搜索
		letfv := float64(x - arr[begin])
		allv := float64(arr[end] - arr[begin])
		diff := float64(end - begin)
		mid = int(float64(begin) + letfv/allv*diff)
		if mid < 0 || mid >= len(arr) {
			return false, -1
		}
		//mid = begin + (end-begin)>>2
		if arr[mid] < x {
			begin = mid + 1
		} else if arr[mid] > x {
			end = mid - 1
		} else {
			return true, mid
		}
	}
	return false, -1
}
