package main

import "math"

func findMin(nums []int) int {
	N := len(nums)
	if N == 0 {
		return math.MaxInt64
	}
	left, mid, right := 0, 0, N-1
	for left < right {
		mid = (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right--
		}
	}
	return nums[left]
}

func main() {

}
