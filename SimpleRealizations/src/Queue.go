package src

type qNode struct {
	key  int
	next *qNode
}
type Queue struct {
	head *qNode
	tail *qNode
	size int
}

func (f *Queue) Enqueue(value int) {
	f.size++
	temp := &qNode{value, nil}

	if f.tail == nil {
		f.head = temp
		f.tail = temp
		return
	}
	f.tail.next = temp
	f.tail = f.tail.next

}
func (f *Queue) Dequeue() (r int) {
	if f.size != 0 {
		r = f.head.key
		f.head = f.head.next
		if f.head == nil {
			f.tail = nil
		}
	} else {
		panic("Queue has no elements")
		return
	}
	f.size--
	return r
}
func (f *Queue) Peek() int {
	if f.size == 0 {
		panic("Queue has no elements")
	}
	return f.head.key
}
func (f *Queue) Count() int {
	return f.size
}
