package main

func findPeakElement(nums []int) int {
	left, mid, right := 0, -1, len(nums)-1
	for left <= right {
		mid = (left + right) / 2
		if mid == 0 || nums[mid] > nums[mid-1] {
			if mid == len(nums)-1 || nums[mid] > nums[mid+1] {
				return mid
			}
		}
		if mid > 0 && nums[mid] < nums[mid-1] {
			right = mid - 1
		} else if mid < len(nums)-1 || nums[mid] < nums[mid+1] {
			left = mid + 1
		}
	}
	return -1
}

func main() {

}
