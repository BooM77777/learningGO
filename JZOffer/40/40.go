package main

import "container/heap"

type MyHeap []int

func (this MyHeap) Len() int {
	return len(this)
}
func (this MyHeap) Less(i, j int) bool {
	return this[i] > this[j]
}
func (this *MyHeap) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}
func (this *MyHeap) Push(data interface{}) {
	(*this) = append((*this), data.(int))
}
func (this *MyHeap) Pop() interface{} {
	ret := (*this)[len(*this)-1]
	(*this) = (*this)[:len(*this)-1]
	return ret
}
func (this *MyHeap) push(data int) {
	heap.Push(this, data)
}
func (this *MyHeap) pop() int {
	return heap.Pop(this).(int)
}
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	h := &MyHeap{}
	for _, data := range arr {
		if h.Len() == k {
			if top := h.pop(); top < data {
				data = top
			}
		}
		h.push(data)
	}
	return *h
}

func main() {

}
