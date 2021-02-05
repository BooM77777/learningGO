package main

import "container/heap"

type Char struct {
	char byte
	freq int
}
type bHeap []Char

func (h bHeap) Less(i, j int) bool {
	return h[i].freq > h[j].freq
}
func (h bHeap) Len() int {
	return len(h)
}
func (h *bHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *bHeap) Push(val interface{}) {
	(*h) = append((*h), val.(Char))
}
func (h *bHeap) Pop() (val interface{}) {
	val = (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return
}
func (h *bHeap) push(val Char) {
	heap.Push(h, val)
}
func (h *bHeap) pop() Char {
	return heap.Pop(h).(Char)
}
func rearrangeString(s string, k int) string {
	if k == 0 {
		return s
	}
	counter := map[byte]int{}
	for i := range s {
		counter[s[i]]++
	}
	h := &bHeap{}
	for k, v := range counter {
		h.push(Char{k, v})
	}
	q := make([]Char, 0, k)
	ret := make([]byte, 0, len(s))
	for len((*h)) > 0 {
		tmpChar := h.pop()
		ret = append(ret, tmpChar.char)
		tmpChar.freq--
		q = append(q, tmpChar)
		if len(q) == k {
			if q[0].freq > 0 {
				h.push(q[0])
			}
			q = q[1:]
		}
	}
	if len(ret) == len(s) {
		return string(ret)
	}
	return ""
}
func main() {

}
