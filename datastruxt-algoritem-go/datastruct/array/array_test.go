package array

import (
	"fmt"
	"testing"
)

// Create by czx on 2019/12/25

func TestArrayLis(t *testing.T) {
	arr := []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	type fields struct {
		element  []interface{}
		size     int
		capacity int
	}
	type args struct {
		arr []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "array list test 1",
			fields: fields{
				element:  []interface{}{},
				size:     0,
				capacity: len(arr),
			},
			args: args{arr: arr},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Array{
				element:  tt.fields.element,
				size:     tt.fields.size,
				capacity: tt.fields.capacity,
			}

			for i := 0; i < len(tt.args.arr); i++ {
				l.Append(tt.args.arr[i])
			}

			l.ForEach(func(v interface{}) {
				fmt.Printf("%d ", v)
			})
			fmt.Println()
			fmt.Println("Size = ", l.Size())
			fmt.Println("capacity: ", l.Capacity())
			if ok, res := l.Pop(0); ok {
				fmt.Println("remove res: ", res)
			}
			if ok, res := l.Pop(l.Size() - 1); ok {
				fmt.Println("res remove: ", res)
			}

			l.ForEach(func(v interface{}) {
				fmt.Printf("%d ", v)
			})

			fmt.Println("Size = ", l.Size())
			fmt.Println("capacity: ", l.Capacity())
			l.Set(5, 100)
			fmt.Println("index 5: ", l.Get(5))

			if ok, _ := l.Pop(100); !ok {
				fmt.Println("remove is not ok")
			}

			fmt.Println("is empty: ", l.Empty())
			l.Insert(4, 500)
			l.ForEach(func(v interface{}) {
				fmt.Printf("%d ", v)
			})

			fmt.Println()
			l.Clear()
			fmt.Println("is empty: ", l.Empty())
			l.ForEach(func(v interface{}) {
				fmt.Printf("%d ", v)
			})

			for i := 0; i < 20; i++ {
				l.Append(i * 10)
			}
			fmt.Println("new capacity: ", l.Capacity())
			fmt.Println("new size: ", l.Size())
			l.ForEach(func(v interface{}) {
				fmt.Printf("%d ", v)
			})
		})
	}
}
