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
			args: args{util.GeneratorArray(10, 1, 10)},
		},
		{
			name: "quick sort 2 way",
			args: args{util.GeneratorArray(20, 1, 10)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			array := util.SortedArray(tt.args.arr)
			Sort2Way(tt.args.arr)
			fmt.Println(tt.name+" result: ", reflect.DeepEqual(array, tt.args.arr))
			fmt.Println(tt.args.arr)
		})
	}
}
