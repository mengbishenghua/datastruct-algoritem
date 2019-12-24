package binarysearch

import (
	"fmt"
	"testing"
)

// Create by czx on 2019/12/24

func TestSearch(t *testing.T) {
	type args struct {
		arr   []int
		begin int
		end   int
		x     int
	}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		{
			name: "binary search 1",
			args: args{
				arr:   arr,
				begin: 0,
				end:   9,
				x:     5,
			},
			want:  true,
			want1: 4,
		},
		{
			name: "binary search 1",
			args: args{
				arr:   arr,
				begin: 0,
				end:   9,
				x:     100,
			},
			want:  false,
			want1: -1,
		},
		{
			name: "binary search 1",
			args: args{
				arr:   arr,
				begin: 0,
				end:   9,
				x:     1,
			},
			want:  true,
			want1: 0,
		},
		{
			name: "binary search 1",
			args: args{
				arr:   arr,
				begin: 0,
				end:   9,
				x:     10,
			},
			want:  true,
			want1: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Search(tt.args.arr, tt.args.begin, tt.args.end, tt.args.x)
			if got != tt.want {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Search() got1 = %v, want %v", got1, tt.want1)
			}
			fmt.Println(got, got1)
		})
	}
}
