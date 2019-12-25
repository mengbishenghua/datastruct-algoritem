package datastruct

// Create by czx on 2019/12/25
type Q interface {
	Push(e interface{})
	PopFront() (bool, interface{})
	Empty() bool
	Peek() interface{}
	Size() int
	Clear()
}
