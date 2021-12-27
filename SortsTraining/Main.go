package main

import (
	Sorts "GoTask1/SortsTraining/src"
)

func main() {
	var sorts Sorts.Sorts
	sorts.Init(6162, 50, 1503)
	sorts.Print()
	sorts.BubbleSort()
	sorts.ShakerSort()
	sorts.Naive()
	sorts.InsertionSort()
	sorts.ComSort()

	sorts.SelectionSort()

	sorts.QuickSort()

	sorts.MergeSort()

	sorts.HeapSort()
}
