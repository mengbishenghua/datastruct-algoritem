package datastruct

// Create by czx on 2019/12/25

type Collection interface {
	Append(e interface{})
	Insert(index int, e interface{})
	Empty() bool
	Size() int
	Get(index int) interface{}
	Set(index int, e interface{})
	pop(index int) (bool, interface{})
	Clear()
}
