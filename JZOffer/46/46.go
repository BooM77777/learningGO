package main

import "strconv"

func translateNum(num int) int {
	str := strconv.Itoa(num)
	dp0, dp1, dp2 := 0, 0, 1
	for i := range str {
		dp0, dp1 = dp1, dp2
		if i == 0 {
			continue
		}
		if str[i-1:i+1] >= "10" && str[i-1:i+1] <= "25" {
			dp2 += dp0
		}
	}
	return dp2
}

func main() {

}
