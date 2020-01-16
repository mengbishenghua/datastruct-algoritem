package thresort

import (
	"encoding/binary"
	"io"
	"math/rand"
	"sort"
	"time"
)

// Create by czx on 2020/1/16

// 多线程归并排序

var start time.Time

// 生成随机数
func GenerateRandom(count int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		for i := 0; i < count; i++ {
			n := rand.Intn(1000)
			out <- n
		}
		close(out)
	}()
	return out
}

// 将要排序的数组分成n份分别进行排序
func ThreadSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		var data []int
		for v := range in {
			data = append(data, v)
		}
		sort.Ints(data)
		for _, v := range data {
			out <- v
		}
		close(out)
	}()
	return out
}

// 将已经排序的两个数组进行归并
func Merge(arr1, arr2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-arr1
		v2, ok2 := <-arr2
		// 将数据升序放入out
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-arr1
			} else {
				out <- v2
				v2, ok2 = <-arr2
			}
		}
		close(out)
	}()
	return out
}

// 递归的进行合并
func MergeSort(arrays ...<-chan int) <-chan int {
	start = time.Now()
	if len(arrays) == 1 {
		return arrays[0]
	}
	m := len(arrays) / 2
	// 将多个管道归并为一个
	return Merge(MergeSort(arrays[:m]...), MergeSort(arrays[:m]...))
}

// 将排好序的数据存储
func WriteSlink(w io.Writer, in <-chan int) {
	for v := range in {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(v))
		w.Write(buf)
	}
}

// 读取数据
func ReadData(r io.Reader, size int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buf := make([]byte, 8)
		readSize := 0
		for {
			n, err := r.Read(buf)
			readSize += n
			if n > 0 {
				out <- int(binary.BigEndian.Uint64(buf))
			}
			if err != nil || (size != -1 && readSize >= size) {
				break
			}
		}
		close(out)
	}()
	return out
}
