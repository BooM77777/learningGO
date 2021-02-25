package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	left, retVal, win := 0, 0, 1
	for right := 0; right < len(nums); right++ {
		win = win * nums[right]
		for ; left <= right && win >= k; left++ {
			win = win / nums[left]
		}
		retVal += right - left + 1
	}
	return retVal
}

func main() {

}
