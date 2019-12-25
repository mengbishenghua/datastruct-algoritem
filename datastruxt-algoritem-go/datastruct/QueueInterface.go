package datastruct

// Create by czx on 2019/12/25
type Q interface {
	Push(e interface{})
	Pop() (bool, interface{})
	Empty() bool
	Peek() interface{}
	Size() int
	Clear()
}
