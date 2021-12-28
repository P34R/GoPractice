package src

type dqNode struct {
	key  int
	next *dqNode
	prev *dqNode
}
type Deque struct {
	front *dqNode
	rear  *dqNode
	size  int
}

func (f *Deque) EnqueueFirst(value int) {
	temp := &dqNode{value, nil, nil}
	if f.front == nil {
		f.front = temp
		f.rear = temp
	} else {
		f.front.prev = temp
		temp.next = f.front
		f.front = temp
	}
	f.size++
}
func (f *Deque) EnqueueLast(value int) {
	temp := &dqNode{value, nil, nil}
	if f.rear == nil {
		f.front = temp
		f.rear = temp
	} else {
		f.rear.next = temp
		temp.prev = f.rear
		f.rear = temp
	}
	f.size++
}
func (f *Deque) PeekFirst() int {
	return f.front.key
}
func (f *Deque) PeekLast() int {
	return f.rear.key
}
func (f *Deque) DequeueFirst() (value int) {
	if f.size == 0 {
		panic("Deque has no elements")
	}
	value = f.front.key
	f.front = f.front.next
	if f.front == nil {
		f.rear = nil
	} else {
		f.front.prev = nil
	}
	f.size--
	return value
}
func (f *Deque) DequeueLast() (value int) {
	if f.size == 0 {
		panic("Deque has no elements")
	}
	value = f.rear.key
	f.rear = f.rear.prev
	if f.rear == nil {
		f.front = nil
	} else {
		f.rear.next = nil
	}
	f.size--
	return value
}
func (f *Deque) Count() int {
	return f.size
}
