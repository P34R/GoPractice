package main

import (
	"GoTask1/SimpleRealizations/src"
	"fmt"
)

func main() {
	var myStack src.Stack
	myStack.Push(1)
	fmt.Println("count - ", myStack.Count())
	myStack.Push(2)
	fmt.Println("count - ", myStack.Count())
	myStack.Push(3)
	fmt.Println("count - ", myStack.Count())
	fmt.Println(myStack.Pop())
	fmt.Println("count - ", myStack.Count())
	fmt.Println(myStack.Pop())
	fmt.Println("count - ", myStack.Count())
	fmt.Println(myStack.Pop())
	fmt.Println("count - ", myStack.Count())
	var myQueue src.Queue
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)
	myQueue.Enqueue(4)
	fmt.Println("size", myQueue.Count())
	fmt.Println("peek", myQueue.Peek())
	fmt.Println("size", myQueue.Count())
	fmt.Println("deq", myQueue.Dequeue())
	fmt.Println("size", myQueue.Count())
	fmt.Println("peek", myQueue.Peek())
	fmt.Println("deq", myQueue.Dequeue())
	fmt.Println("deq", myQueue.Dequeue())
	fmt.Println("deq", myQueue.Dequeue())
	var myDeque src.Deque
	myDeque.EnqueueFirst(1)
	myDeque.EnqueueFirst(2)
	myDeque.EnqueueFirst(3)
	myDeque.EnqueueFirst(4)
	myDeque.EnqueueFirst(5)
	myDeque.EnqueueLast(6)
	fmt.Println("DEQ FIRST: ", myDeque.DequeueFirst())
	fmt.Println("SIZE: ", myDeque.Count())
	fmt.Println("DEQ LAST: ", myDeque.DequeueLast())
	fmt.Println("SIZE: ", myDeque.Count())
	fmt.Println("DEQ FIRST: ", myDeque.DequeueFirst())
	fmt.Println("SIZE: ", myDeque.Count())
	fmt.Println("DEQ LAST: ", myDeque.DequeueLast())
	fmt.Println("SIZE: ", myDeque.Count())
	fmt.Println("DEQ FIRST: ", myDeque.DequeueFirst())
	fmt.Println("SIZE: ", myDeque.Count())
	fmt.Println("DEQ FIRST: ", myDeque.DequeueFirst())
	fmt.Println("SIZE: ", myDeque.Count())

}
