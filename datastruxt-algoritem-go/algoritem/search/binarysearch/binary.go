package binarysearch

// Create by czx on 2019/12/24

//Search 二分搜索
func Search(arr []int, begin, end, x int) (bool, int) {
	if len(arr) == 0 {
		return false, -1
	}
	var mid = 0
	for begin <= end {
		mid = begin + (end-begin)>>2
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
