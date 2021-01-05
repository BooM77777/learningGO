package main

import (
	"fmt"
)

//MinMaxHeap 最小最大堆
type MinMaxHeap struct {
	heap []int
}

//NewMinMaxHeap 创建一个空的最小最大堆
func NewMinMaxHeap() *MinMaxHeap {
	ret := &MinMaxHeap{make([]int, 1)} // 为了方便计算，设置一个哑节点
	return ret
}

func (h *MinMaxHeap) getLayer(idx int) (layer int) {

	for idx > 0 {
		idx = idx >> 1
		layer++
	}
	layer = layer - 1
	return
}

func (h *MinMaxHeap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *MinMaxHeap) min(input ...int) (idx int) {
	idx = input[0]
	for i := 1; i < len(input); i++ {
		if input[i] > 0 && input[i] <= h.heap[0] && h.heap[input[i]] < h.heap[idx] {
			idx = input[i]
		}
	}
	return
}

func (h *MinMaxHeap) max(input ...int) (idx int) {
	idx = input[0]
	for i := 1; i < len(input); i++ {
		if input[i] > 0 && input[i] <= h.heap[0] && h.heap[input[i]] > h.heap[idx] {
			idx = input[i]
		}
	}
	return
}

func (h *MinMaxHeap) fixUpFromDownToUp(idx int) {

	layer := h.getLayer(idx)

	var pIdx int

	if layer%2 == 0 {
		// 如果是处在小层
		for {
			pIdx = idx / 4
			if idx <= 1 || h.heap[pIdx] <= h.heap[idx] {
				break
			}
			h.swap(pIdx, idx)
			idx = pIdx
		}
	} else {
		// 如果处在大层
		for {
			pIdx = idx / 4
			if idx <= 3 || h.heap[pIdx] >= h.heap[idx] {
				break
			}
			h.swap(pIdx, idx)
			idx = pIdx
		}
	}
}

func (h *MinMaxHeap) fixUpFromUpToDown(idx int) int {

	layer := h.getLayer(idx)
	var cIdx int

	if layer%2 == 0 {
		for {
			cIdx = h.min(idx, idx*4, idx*4+1, idx*4+2, idx*4+3)
			// fmt.Println(cIdx, idx, idx*4, idx*4+1, idx*4+2, idx*4+3)
			if cIdx == idx {
				break
			}
			h.swap(cIdx, idx)
			idx = cIdx
		}
	} else {
		for {
			cIdx = h.max(idx, idx*4, idx*4+1, idx*4+2, idx*4+3)
			if cIdx == idx {
				break
			}
			h.swap(cIdx, idx)
			idx = cIdx
		}
	}
	return idx
}

//Insert 最小最大堆的节点插入
func (h *MinMaxHeap) Insert(data int) {

	h.heap = append(h.heap, data) // 将节点插入至尾部
	h.heap[0]++                   // 调整计数器
	idx := h.heap[0]

	if h.heap[0] <= 1 {
		return
	}

	layer := h.getLayer(idx) // 获取当前层数
	// 由于是新插入的节点，因此必不会有子节点
	if layer%2 == 0 {
		// 如果处于小层，但是比父节点大\
		if h.max(idx, idx/2) == idx {
			h.swap(idx, idx/2)
			h.fixUpFromDownToUp(idx / 2)
		}
		h.fixUpFromDownToUp(idx)
	} else {
		//如果处于大层，但是比父节点小
		if h.min(idx, idx/2) == idx {
			h.swap(idx, idx/2)
			h.fixUpFromDownToUp(idx / 2)
		}
		h.fixUpFromDownToUp(idx)
	}
}

//PopMin 弹出最小值
func (h *MinMaxHeap) PopMin() (ret int, ok bool) {

	if h.heap[0] == 0 {
		return 0, false
	}

	ret = h.heap[1]
	h.swap(1, h.heap[0])
	h.heap = h.heap[:h.heap[0]] // 有一个哑节点
	h.heap[0]--

	// 首先自顶向下维护小堆
	idx := h.fixUpFromUpToDown(1)
	// fmt.Println(h.heap)

	layer := h.getLayer(idx)

	if layer%2 == 0 {
		// 当前在小堆
		// 与父节点、子节点进行比较
		tmpIdx := h.min(idx, idx/2, idx*2, idx*2+1)
		if tmpIdx != idx {
			h.swap(tmpIdx, idx)
			h.fixUpFromDownToUp(tmpIdx)
		}
	} else {
		tmpIdx := h.max(idx, idx/2, idx*2, idx*2+1)
		if tmpIdx != idx {
			h.swap(tmpIdx, idx)
			h.fixUpFromDownToUp(tmpIdx)
		}
	}
	return ret, true
}

//PopMax 弹出最大值
func (h *MinMaxHeap) PopMax() (ret int, ok bool) {

	if h.heap[0] == 0 {
		return 0, false
	}

	idx := h.max(1, 2, 3)
	ret = h.heap[idx]
	h.swap(idx, h.heap[0])
	h.heap = h.heap[:h.heap[0]] // 有一个哑节点
	h.heap[0]--

	// 首先自顶向下维护堆
	idx = h.fixUpFromUpToDown(idx)
	if idx > h.heap[0] {
		return ret, true
	}
	// fmt.Println(h.heap)

	layer := h.getLayer(idx)

	if layer%2 == 0 {
		// 当前在小堆
		// 与父节点、子节点进行比较
		tmpIdx := h.min(idx, idx/2, idx*2, idx*2+1)
		if tmpIdx != idx {
			h.swap(tmpIdx, idx)
			h.fixUpFromDownToUp(tmpIdx)
		}
	} else {
		tmpIdx := h.max(idx, idx/2, idx*2, idx*2+1)
		if tmpIdx != idx {
			h.swap(tmpIdx, idx)
			h.fixUpFromDownToUp(tmpIdx)
		}
	}
	return ret, true
}

func main() {
	heap := NewMinMaxHeap()
	input := []int{12, 60, 75, 23, 39, 58, 21, 42, 30, 65, 18, 20, 10, 15, 75}
	for _, item := range input {
		heap.Insert(item)
	}
	for i := 0; i < 20; i++ {
		fmt.Println(heap.PopMax())
	}
}
