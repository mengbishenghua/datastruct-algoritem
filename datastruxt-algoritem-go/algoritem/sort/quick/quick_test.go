package quick

import (
	"com.datastruct.algoritem.go/algoritem/util"
	"fmt"
	"reflect"
	"testing"
)

// Create by czx on 2019/12/31

func TestSort2Way(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "quick sort 2 way",
			args: args{util.GeneratorArray(100, 1, 100)},
		},
		{
			name: "quick sort 2 way",
			args: args{util.GeneratorArray(10000, 1, 1000)},
		},
		{
			name: "quick sort 2 way",
			args: args{util.GeneratorArray(1000000, 1, 10000)},
		},
		{
			name: "quick sort 2 way",
			args: args{util.GeneratorArray(10000000, 1, 10000)},
		},
		{
			name: "quick sort 3 way",
			args: args{util.GeneratorArray(100, 1, 100)},
		},
		{
			name: "quick sort 3 way",
			args: args{util.GeneratorArray(10000, 1, 1000)},
		},
		{
			name: "quick sort 3 way",
			args: args{util.GeneratorArray(1000000, 1, 10000)},
		},
		{
			name: "quick sort 3 way",
			args: args{util.GeneratorArray(10000000, 1, 10000)},
		},
	}
	for i, tt := range tests {
		if i <= 3 {
			t.Run(tt.name, func(t *testing.T) {
				array := util.SortedArray(tt.args.arr)
				Sort2Way(tt.args.arr)
				fmt.Println(tt.name+" result: ", reflect.DeepEqual(array, tt.args.arr))
			})
		} else {
			t.Run(tt.name, func(t *testing.T) {
				array := util.SortedArray(tt.args.arr)
				Sort3way(tt.args.arr)
				fmt.Println(tt.name+" result: ", reflect.DeepEqual(array, tt.args.arr))
			})
		}
	}
}
