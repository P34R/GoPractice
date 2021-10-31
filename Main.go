package main

import (
	"GoTask1/src"
	"fmt"
)

func main() {
	/*	var myStack src.Stack
		myStack.Push(1)
		fmt.Println(myStack.Count())
		fmt.Println(myStack.Peek())
		fmt.Println(myStack.Pop())
		fmt.Println(myStack.Pop())
		fmt.Println(myStack.Count())*/
	var myQueue src.Queue
	myQueue.Enqueue(5)
	myQueue.Enqueue(3)
	myQueue.Enqueue(1)
	myQueue.Enqueue(4)
	fmt.Println(myQueue.Count())
	fmt.Println(myQueue.Peek())
	fmt.Println(myQueue.Count())
	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue.Count())
	fmt.Println(myQueue.Peek())
}
