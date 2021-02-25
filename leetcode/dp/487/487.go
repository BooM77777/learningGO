package main

func findMaxConsecutiveOnes(nums []int) int {
	dp1, dp0, ret := 0, 0, 0
	for _, num := range nums {
		if num == 1 {
			dp0, dp1 = dp0+1, dp1+1
		} else {
			dp0, dp1 = 0, dp0+1
		}
		if ret < dp0 {
			ret = dp0
		}
		if ret < dp1 {
			ret = dp1
		}
	}
	return ret
}
func main() {

}
