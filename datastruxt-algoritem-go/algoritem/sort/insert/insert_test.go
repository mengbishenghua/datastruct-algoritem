package insert

import (
	"com.datastruct.algoritem.go/algoritem/util"
	"fmt"
	"reflect"
	"testing"
)

// Create by czx on 2019/12/24

func TestSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"insert sort1",
			args{arr: util.GeneratorArray(100, 10, 100)},
		},
		{
			"insert sort2",
			args{arr: util.GeneratorArray(1000, 100, 1000)},
		},
		{
			"insert sort3",
			args{arr: util.GeneratorArray(10000, 100, 1000)},
		},
		{
			"insert sort4",
			args{arr: util.GeneratorArray(100000, 1000, 10000)},
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
