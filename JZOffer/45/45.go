package main

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		str1, str2 := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		str12, str21 := str1+str2, str2+str1
		num12, _ := strconv.Atoi(str12)
		num21, _ := strconv.Atoi(str21)
		return num12 < num21
	})
	ret := ""
	for i := range nums {
		ret += strconv.Itoa(nums[i])
	}
	return ret
}

func main() {

}
