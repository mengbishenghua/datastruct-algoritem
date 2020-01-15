package heap

import (
	"com.datastruct.algoritem.go/algoritem/util"
	"fmt"
	"reflect"
	"testing"
)

// Create by czx on 2020/1/1

func TestSOrt(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "heap sort",
			args: args{util.GeneratorArray(10, 1, 10)},
		},
		{
			name: "heap sort",
			args: args{util.GeneratorArray(1000, 10, 100)},
		},
		{
			name: "heap sort",
			args: args{util.GeneratorArray(100000, 10, 1000)},
		},
		{
			name: "heap sort",
			args: args{util.GeneratorArray(10000000, 100, 10000)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			array := util.SortedArray(tt.args.arr)
			Sort(tt.args.arr)
			fmt.Println(tt.name+" result: ", reflect.DeepEqual(array, tt.args.arr))
		})
	}
}
