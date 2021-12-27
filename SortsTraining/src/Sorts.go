package Sorts

import (
	"fmt"
	"math/rand"
	"sort"
)

type Sorts struct {
	unsorted []int
	data     []int
	sorted   []int
}

func (s *Sorts) Init(key int64, size int, max int) {
	s.data = make([]int, size)
	s.unsorted = make([]int, size)
	s.sorted = make([]int, size)
	if max < len(s.data) {
		panic("Too low maximum number\n")
	}
	rand.Seed(key)
	for i := 0; i < len(s.data); i++ {
		s.data[i] = rand.Int() % max

	}
	copy(s.unsorted, s.data)
	copy(s.sorted, s.data)
	sort.Ints(s.sorted)
}
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func (s *Sorts) Print() {
	fmt.Println(s.data)
}
func (s *Sorts) swap(i int, j int) {
	swap := s.data[i]
	s.data[i] = s.data[j]
	s.data[j] = swap
}
func (s *Sorts) check(message string) {
	if equal(s.data, s.sorted) {
		fmt.Println(message, " - OK")
	}
}
func (s *Sorts) Naive() {
	copy(s.data, s.unsorted)
	for i := 0; i < len(s.data); i++ {
		for j := i + 1; j < len(s.data); j++ {
			if s.data[i] > s.data[j] {
				s.swap(i, j)
			}
		}
	}
	s.check("Naive")
}
func (s *Sorts) BubbleSort() {
	copy(s.data, s.unsorted)
	for i := 0; i < len(s.data)-1; i++ {
		for j := 0; j < len(s.data)-1; j++ {
			if s.data[j] > s.data[j+1] {
				s.swap(j, j+1)
			}
		}
	}
	s.check("Bubble")
}
func (s *Sorts) ShakerSort() {
	copy(s.data, s.unsorted)
	left := 0
	right := len(s.data) - 1
	for left < right {
		for i := left; i < right; i++ {
			if s.data[i] > s.data[i+1] {
				s.swap(i, i+1)
			}
		}
		right--
		for i := right; i > left; i-- {
			if s.data[i] < s.data[i-1] {
				s.swap(i, i-1)
			}
		}
		left++
	}
	s.check("Shaker")
}
func (s *Sorts) ComSort() {
	copy(s.data, s.unsorted)
	factor := 1.247
	step := len(s.data) - 1
	for step > 1 {
		for i := 0; i < len(s.data)-step; i++ {
			if s.data[i] > s.data[i+1] {
				s.swap(i, i+1)
			}
		}
		step = int(float64(step) / factor)
	}
	for i := 0; i < len(s.data)-1; i++ {
		isSwap := false
		for j := 0; j < len(s.data)-1; j++ {
			if s.data[j] > s.data[j+1] {
				isSwap = true
				s.swap(j, j+1)
			}
		}
		if !isSwap {
			break
		}
	}
	s.check("Comb")
}
func (s *Sorts) InsertionSort() {
	copy(s.data, s.unsorted)
	for i := 1; i < len(s.data); i++ {
		x := s.data[i]
		j := i
		for j > 0 && s.data[j-1] > x {
			s.data[j] = s.data[j-1]
			j--
		}
		s.data[j] = x
	}
	s.check("Insertion")
}
func (s *Sorts) minimum(i, j int) int {
	if s.data[i] > s.data[j] {
		return j
	}
	return i
}
func (s *Sorts) SelectionSort() {
	copy(s.data, s.unsorted)
	for i := 0; i < len(s.data)-1; i++ {
		min := i
		for j := i + 1; j < len(s.data); j++ {
			if s.data[j] < s.data[min] {
				min = j
			}
		}
		s.swap(i, min)
	}
	s.check("Selection")
}

func (s *Sorts) partition(l, r int) int {
	x := s.data[r]
	less := l
	for i := l; i < r; i++ {
		if s.data[i] <= x {
			s.swap(i, less)
			less++
		}
	}
	s.swap(less, r)
	return less
}

func (s *Sorts) quickSort(l, r int) {
	if l < r {
		pivot := s.partition(l, r)
		s.quickSort(l, pivot-1)
		s.quickSort(pivot+1, r)
	}
}
func (s *Sorts) QuickSort() {
	copy(s.data, s.unsorted)
	s.quickSort(0, len(s.data)-1)
	s.check("QuickSort")
}

func (s *Sorts) mergeSort(buffer []int, l int, r int) {
	if l < r {
		var m int
		m = (l + r) / 2
		s.mergeSort(buffer, l, m)
		s.mergeSort(buffer, m+1, r)
		k := l
		for i, j := l, m+1; i <= m || j <= r; {
			if j > r || (i <= m && s.data[i] < s.data[j]) {
				buffer[k] = s.data[i]
				i++
			} else {
				buffer[k] = s.data[j]
				j++
			}
			k++
		}
		for i := l; i <= r; i++ {
			s.data[i] = buffer[i]
		}
	}
}

func (s *Sorts) MergeSort() {
	copy(s.data, s.unsorted)
	buffer := make([]int, len(s.data))
	s.mergeSort(buffer, 0, len(s.data)-1)
	s.check("MergeSort")
}

func (s *Sorts) heapify(n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && s.data[l] > s.data[largest] {
		largest = l
	}
	if r < n && s.data[r] > s.data[largest] {
		largest = r
	}
	if largest != i {
		s.swap(i, largest)
		s.heapify(n, largest)
	}

}
func (s *Sorts) HeapSort() {
	copy(s.data, s.unsorted)
	for i := len(s.data)/2 - 1; i >= 0; i-- {
		s.heapify(len(s.data), i)
	}
	for i := len(s.data) - 1; i >= 0; i-- {
		s.swap(0, i)
		s.heapify(i, 0)
	}
	s.check("HeapSort")
}
