package main

import "math/rand"

type Solution struct {
	sum int
	w   []int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}

	s := Solution{sum: w[len(w)-1], w: make([]int, len(w)+1)}
	copy(s.w[1:], w)
	return s
}

func (this *Solution) PickIndex() int {
	p := rand.Intn(this.sum)
	left, mid, right := 1, -1, len(this.w)-1
	for left <= right {
		// fmt.Println(left, right)
		mid = (left + right) / 2
		if this.w[mid-1] <= p && p < this.w[mid] {
			return mid - 1
		} else if this.w[mid-1] > p {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
func main() {

}
