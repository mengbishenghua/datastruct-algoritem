package thresort

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

// Create by czx on 2020/1/16

func TestMerge(t *testing.T) {
	var filename = "data.in"
	var count = 100
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pipeline := GenerateRandom(count)
	w := bufio.NewWriter(file)
	WriteSlink(w, pipeline)
	w.Flush()
	fmt.Println("生成数据")
	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("读取数据")
	//data := ReadData(bufio.NewReader(file), -1)
	//for v := range data {
	//	//fmt.Println(v)
	//}
}

// 测试2亿数据：耗时：163.5626ms
func TestMergeSort(t *testing.T) {
	out := MergeSort(
		ThreadSort(GenerateRandom(100000)),
		ThreadSort(GenerateRandom(100000)),
	)
	fmt.Println("耗时: ", time.Since(start))
	//time.Sleep(5 * time.Second)
	count := 0
	for v := range out {
		count += 1
		fmt.Println("v: ", v, " count: ", count)
	}
}

func TestSort(t *testing.T) {
	out := Merge(
		ThreadSort(ArraySource(3, 8, 1, 10, 4)),
		ThreadSort(ArraySource(13, 18, 11, 110, 14)),
	)

	for v := range out {
		fmt.Println(v)
	}
}

func ArraySource(inputs ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range inputs {
			out <- v
		}
		close(out)
	}()
	return out
}
