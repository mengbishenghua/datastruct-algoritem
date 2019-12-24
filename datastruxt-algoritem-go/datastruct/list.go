package datastruct

// Create by czx on 2019/12/25

type List interface {
	Append(e interface{})
	Insert(index int, e interface{})
	Empty() bool
	Size() int
	Capacity() int
	Get(index int) interface{}
	Set(index int, e interface{})
	Remove(index int) (bool, interface{})
	Clear()
}
