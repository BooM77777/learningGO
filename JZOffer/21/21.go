package main

func exchange(nums []int) []int {
	iterFast, iterSlow := 0, 0
	for iterFast = range nums {
		if nums[iterFast]%2 == 0 {
			continue
		}
		nums[iterSlow], nums[iterFast] = nums[iterFast], nums[iterSlow]
		iterSlow++
	}
	return nums
}

func main() {

}
