package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) (ret int) {
	N, k, j := len(nums), 0, 0
	if N == 0 {
		return 0
	}
	hash := map[int]bool{}
	for _, num := range nums {
		hash[num] = true
	}
	for k = range hash {
		if _, ok := hash[k-1]; !ok {
			for j = 1; hash[k+j]; j++ {
			}
			ret = max(ret, j)
		}
	}
	return
}

func main() {

}
