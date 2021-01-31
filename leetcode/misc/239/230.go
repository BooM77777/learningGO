package main

func insert(val int, win *[]int) {
	for len(*win) > 0 {
		if val > (*win)[len(*win)-1] {
			(*win) = (*win)[:len(*win)-1]
		} else {
			break
		}
	}
	(*win) = append((*win), val)
}
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k {
		//
	}
	win, ret := []int{}, make([]int, 0, len(nums)-k)
	var i int
	for i = 0; i < k; i++ {
		insert(nums[i], &win)
	}
	ret = append(ret, win[0])
	for ; i < len(nums); i++ {
		if nums[i-k] == win[0] {
			win = win[1:]
		}
		insert(nums[i], &win)
		ret = append(ret, win[0])
	}
	return ret
}
func main() {

}
