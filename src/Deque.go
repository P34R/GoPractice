package src

type Deque struct {
	list []int
}

func (f *Deque) EnqueueFirst(value int) {
	f.list = append(f.list, value)
	f.list = append(f.list[len(f.list)-1:len(f.list)], f.list[:len(f.list)-1]...)
}
func (f *Deque) EnqueueLast(value int) {
	f.list = append(f.list, value)
}
func (f *Deque) PeekFirst() int {
	if len(f.list) == 0 {
		return 0 //error, deque empty
	}
	return f.list[0]
}
func (f *Deque) PeekLast() int {
	if len(f.list) == 0 {
		return 0 //error, deque empty
	}
	return f.list[len(f.list)-1]
}
func (f *Deque) DequeueFirst() int {
	if len(f.list) == 0 {
		return 0 //error, deque empty
	}
	first := f.list[0]
	f.list = append(f.list[1:])
	return first
}
func (f *Deque) DequeueLast() int {
	if len(f.list) == 0 {
		return 0 //error, deque empty
	}
	last := f.list[len(f.list)-1]
	f.list = append(f.list[:len(f.list)-1])
	return last
}
func (f *Deque) Count() int {
	return len(f.list)
}
