package main

import "fmt"

// Create by czx on 2019/12/25

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr = append(arr[:4], arr[3:]...)
	arr[3] = 100
	fmt.Println(arr)
	index := 1
	copy(arr[index:], arr[index+1:])
	fmt.Println(arr[:len(arr)-1])
}
