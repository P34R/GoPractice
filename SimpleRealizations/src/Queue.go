package src

type Queue struct {
	list     []int
	size     int
	capacity int
}

func (f *Queue) Enqueue(value int) {
	f.list = append(f.list, value)
	f.list = append(f.list[len(f.list)-1:len(f.list)], f.list[:len(f.list)-1]...)
}
func (f *Queue) Dequeue() int {
	if len(f.list) == 0 {
		return 0 //error, empty queue
	}
	last := f.list[len(f.list)-1]
	f.list = append(f.list[0 : len(f.list)-1])
	return last
}
func (f *Queue) Peek() int {
	if len(f.list) == 0 {
		return 0 //error, empty queue
	}
	return f.list[len(f.list)-1]
}
func (f *Queue) Count() int {
	return len(f.list)
}
