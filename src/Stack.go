package src

type Stack struct {
	list []int
}

func (f *Stack) Push(value int) {
	f.list = append(f.list, value)
}
func (f *Stack) Pop() int {
	if len(f.list) == 0 {
		return -1
	}
	var lastElement int
	lastElement = f.list[len(f.list)-1]
	f.list = append(f.list[:len(f.list)-1], f.list[len(f.list):]...)
	return lastElement
}
func (f *Stack) Peek() int {
	if len(f.list) == 0 {
		return -1
	}
	return f.list[len(f.list)-1]
}
func (f *Stack) Count() int {
	return len(f.list)
}
