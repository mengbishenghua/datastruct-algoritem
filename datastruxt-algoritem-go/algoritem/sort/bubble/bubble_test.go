package bubble

import (
	"com.datastruct.algoritem.go/algoritem/util"
	"fmt"
	"reflect"
	"testing"
)

// Create by czx on 2019/12/24

func TestBubbleSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"bubble sort 1",
			args{[]int{62, 98, 13, 345, 213, 42, 74, 25, 74, 84}},
		},
		{
			"bubble sort 2",
			args{util.GeneratorArray(10, 50, 100)},
		},
		{
			"bubble sort 3",
			args{util.GeneratorArray(10000, 10, 1000)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortedArray := util.SortedArray(tt.args.arr)
			Sort(tt.args.arr)
			fmt.Println(tt.name+":\t", reflect.DeepEqual(sortedArray, tt.args.arr))
		})
	}
}
