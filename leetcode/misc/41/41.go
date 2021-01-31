package main

func firstMissingPositive(nums []int) int {
	N := len(nums)
	for i := range nums {
		for nums[i] > 0 && nums[i] <= N && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return N + 1
}
func main() {

}
