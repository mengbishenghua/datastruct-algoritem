package util

import (
	"math/rand"
	"sort"
	"time"
)

// Create by czx on 2019/12/24

//GeneratorArray 生成随机数组
func GeneratorArray(n, begin, end int) []int {
	arr := make([]int, 0, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(end-begin)+begin)
	}
	return arr
}

//SortedArray 基于arr生成一份新的数组，用于test
func SortedArray(arr []int) []int {
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = arr[i]
	}
	sort.Ints(res)
	return res
}

//Swap 交换
func Swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
