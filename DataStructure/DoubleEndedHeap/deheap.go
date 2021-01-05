package main

import "fmt"

//DEHeap 双端堆，完成了
type DEHeap struct {
	heap []int
}

//NewDEHeap 构造函数
func NewDEHeap() *DEHeap {
	return &DEHeap{make([]int, 1)} // 保留两个哑节点方便计算，实际上只要一个
}

func (h *DEHeap) getLayer(idx int) (layer int) {

	tmp := idx + 1
	for tmp > 1 {
		tmp = tmp >> 1
		layer++
	}
	return
}

func (h *DEHeap) isLeft(idx int) (isLeft bool) {
	// 查询idx在左子树还是右子树
	layer := h.getLayer(idx)

	nodePreSide := 1 << (layer - 1)
	beginNode := nodePreSide << 1
	if idx+1 < beginNode+nodePreSide {
		isLeft = true
	}
	return
}

//Insert 插入值
func (h *DEHeap) Insert(data int) {
	h.heap = append(h.heap, data)
	h.heap[0]++
	if h.heap[0] == 1 {
		return
	}
	idx := h.heap[0]
	if h.isLeft(idx) {
		// idx在左子树
		idx, flag := h.swapFromLeftToRight(idx)
		if flag {
			h.maxHeapFixDownToUp(idx)
		} else {
			h.minHeapFixDownToUp(idx)
		}
	} else {
		// idx在右子树
		idx, flag := h.swapFromRightToLeft(idx)
		if flag {
			// 由于左子树的节点严格小于对应右子树的节点
			// 而relatedIdx在插入前的对应节点对待插入节点的父节点
			// 因此，relatedIdx在交换到右子树后，一定会小于其父节点
			// 所以，只需要对idx节点及左子树进行小堆的维护
			h.minHeapFixDownToUp(idx)
		} else {
			h.maxHeapFixDownToUp(idx)
		}
	}
}

func (h *DEHeap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *DEHeap) min(idxList ...int) (retIdx int) {
	retIdx = idxList[0]
	for i := 1; i < len(idxList); i++ {
		if idxList[i] <= h.heap[0] && h.heap[idxList[i]] < h.heap[retIdx] {
			retIdx = idxList[i]
		}
	}
	return
}

func (h *DEHeap) max(idxList ...int) (retIdx int) {
	retIdx = idxList[0]
	for i := 1; i < len(idxList); i++ {
		if idxList[i] <= h.heap[0] && h.heap[idxList[i]] > h.heap[retIdx] {
			retIdx = idxList[i]
		}
	}
	return
}

func (h *DEHeap) minHeapFixUpToDown(idx int) int {
	for {
		cIdx := idx*2 + 1
		tmpIdx := h.min(idx, cIdx, cIdx+1)
		if tmpIdx == idx {
			break
		}
		h.swap(idx, tmpIdx)
		idx = tmpIdx
	}
	return idx
}

func (h *DEHeap) maxHeapFixUpToDown(idx int) int {
	for {
		cIdx := idx*2 + 1
		tmpIdx := h.max(idx, cIdx, cIdx+1)
		if tmpIdx == idx {
			break
		}
		h.swap(idx, tmpIdx)
		idx = tmpIdx
	}
	return idx
}

func (h *DEHeap) minHeapFixDownToUp(idx int) int {
	for idx > 1 {
		pIdx := (idx - 1) / 2
		if h.heap[idx] >= h.heap[pIdx] {
			break
		}
		h.heap[idx], h.heap[pIdx] = h.heap[pIdx], h.heap[idx]
		idx = pIdx
	}
	return idx
}

func (h *DEHeap) maxHeapFixDownToUp(idx int) int {
	for idx > 2 {
		pIdx := (idx - 1) / 2
		if h.heap[idx] <= h.heap[pIdx] {
			break
		}
		h.heap[idx], h.heap[pIdx] = h.heap[pIdx], h.heap[idx]
		idx = pIdx
	}
	return idx
}

func (h *DEHeap) swapFromLeftToRight(idx int) (int, bool) {

	layer := h.getLayer(idx)

	rIdx := idx + 1<<(layer-1)
	for rIdx > h.heap[0] {
		rIdx = (rIdx - 1) / 2
	}

	if rIdx <= h.heap[0] && h.heap[rIdx] < h.heap[idx] {
		h.swap(rIdx, idx)
		return rIdx, true
	}
	return idx, false
}

func (h *DEHeap) swapFromRightToLeft(idx int) (int, bool) {

	layer := h.getLayer(idx)

	lIdx := idx - 1<<(layer-1)
	lIdx = h.max(lIdx, lIdx*2+1, lIdx*2+2)
	if h.heap[lIdx] > h.heap[idx] {
		h.swap(lIdx, idx)
		return lIdx, true
	}
	return idx, false
}

//PopMin 弹出最小值
func (h *DEHeap) PopMin() (ret int, isEmpty bool) {

	if h.heap[0] == 0 {
		return 0, true
	}

	isEmpty = false

	// 交换最小值和最后一个数，然后弹出
	h.heap[1], h.heap[h.heap[0]] = h.heap[h.heap[0]], h.heap[1]
	ret = h.heap[h.heap[0]]
	h.heap[0]--
	h.heap = h.heap[:h.heap[0]+1]

	idx := h.minHeapFixUpToDown(1)
	if h.heap[0] > 1 {
		idx, flag := h.swapFromLeftToRight(idx)
		if flag {
			h.maxHeapFixDownToUp(idx)
		}
	}

	return
}

//PopMax 弹出最大值
func (h *DEHeap) PopMax() (ret int, isEmpty bool) {

	if h.heap[0] == 0 {
		// 考虑弹出空值的情况
		return 0, true
	}

	// 考虑最大值下标问题，当只有一个值时，最大值下标为1
	maxValIdx := 2
	if h.heap[0] == 1 {
		maxValIdx = 1
	}
	// 交换最大值和最后一个数，然后弹出
	h.heap[maxValIdx], h.heap[h.heap[0]] = h.heap[h.heap[0]], h.heap[maxValIdx]
	ret = h.heap[h.heap[0]]
	h.heap[0]--
	h.heap = h.heap[:h.heap[0]+1]

	if h.heap[0] >= 2 {
		idx := h.maxHeapFixUpToDown(2)
		idx, flag := h.swapFromRightToLeft(idx)
		if flag {
			h.minHeapFixDownToUp(idx)
		}
	}

	return
}

func main() {

	intputSeq := []int{5, 12, 21, 25, 27, 30, 22, 29, 18, 16, 35, 32}
	heap := NewDEHeap()

	for _, item := range intputSeq {
		heap.Insert(item)
	}

	fmt.Println(heap.heap[1:])

	for i := 0; i < len(intputSeq)+5; i++ {
		fmt.Println(heap.PopMin())
		fmt.Println(heap.heap[1:])
	}
}
