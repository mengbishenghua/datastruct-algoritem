package oddeven

import (
	"com.datastruct.algoritem.go/algoritem/util"
	"fmt"
	"reflect"
	"testing"
)

// Create by czx on 2020/1/15

func TestOddEvenSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "odd even sort 1",
			args: args{util.GeneratorArray(100, 1, 100)},
		},
		{
			name: "odd even sort 2",
			args: args{util.GeneratorArray(1000, 10, 1000)},
		},
		{
			name: "odd even sort 3",
			args: args{util.GeneratorArray(10000, 10, 1000)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			array := util.SortedArray(tt.args.arr)
			Sort(tt.args.arr)
			fmt.Println(tt.name+": ", reflect.DeepEqual(array, tt.args.arr))
		})
	}
}
