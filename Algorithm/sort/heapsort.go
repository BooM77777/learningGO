package main

import (
	"container/heap"
	"fmt"
)

// MinHeap ...
type MinHeap []int

// Less ...
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Len ...
func (h MinHeap) Len() int {
	return len(h)
}

// Swap ...
func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Push ...
func (h *MinHeap) Push(data interface{}) {
	(*h) = append((*h), data.(int))
}

// Pop ...
func (h *MinHeap) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return ret
}

func (h *MinHeap) push(data int) {
	heap.Push(h, data)
}

func (h *MinHeap) pop() int {
	return heap.Pop(h).(int)
}

// HeapSort ...
func HeapSort(nums *[]int) {
	h := &MinHeap{}
	for _, num := range *nums {
		h.push(num)
	}
	for i := range *nums {
		(*nums)[i] = h.pop()
	}
}

func main() {
	nums := []int{5, 7, 3, 4, 6, 8, 9, 10, 11, 4, 2}
	HeapSort(&nums)
	fmt.Println(nums)
}
