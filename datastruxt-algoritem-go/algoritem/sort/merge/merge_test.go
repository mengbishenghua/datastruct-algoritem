package merge

import (
	"com.datastruct.algoritem.go/algoritem/util"
	"fmt"
	"reflect"
	"testing"
)

// Create by czx on 2019/12/30

func TestMerge(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "merge sort1",
			args: args{util.GeneratorArray(100, 1, 100)},
		},
		{
			name: "merge sort2",
			args: args{util.GeneratorArray(1000, 1, 100)},
		},
		{
			name: "merge sort3",
			args: args{util.GeneratorArray(10000, 1, 100)},
		},
		{
			name: "merge sort4",
			args: args{util.GeneratorArray(100000, 1, 100)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			array := util.SortedArray(tt.args.arr)
			Merge(tt.args.arr)
			fmt.Println(tt.name+" result: ", reflect.DeepEqual(tt.args.arr, array))
		})
	}
}
