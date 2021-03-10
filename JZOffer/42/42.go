package main

func maxSubArray(nums []int) (retVal int) {
	retVal = -100
	dp := 0
	for i := range nums {
		if dp < 0 {
			dp = nums[i]
		} else {
			dp += nums[i]
		}
		if retVal < dp {
			retVal = dp
		}
	}
	return
}

func main() {

}
