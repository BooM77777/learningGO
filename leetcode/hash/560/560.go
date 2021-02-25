package main

func subarraySum(nums []int, k int) int {
	retVal, pre, hash := 0, 0, map[int]int{}
	hash[0] = 1
	for _, num := range nums {
		pre += num
		if cnt, ok := hash[pre-k]; ok {
			retVal += cnt
		}
		hash[pre]++
	}
	return retVal
}
func main() {

}
