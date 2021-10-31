package main

import (
	"GoTask1/src"
	"fmt"
)

func main() {
	var myStack src.Stack
	myStack.Push(1)
	fmt.Println(myStack.Count())
	fmt.Println(myStack.Peek())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Count())
}
