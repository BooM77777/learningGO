package main

func missingNumber(nums []int) int {
	for i := range nums {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

func main() {

}
