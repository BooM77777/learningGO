package main

func minWindow(s string, t string) (res string) {
	dp, min := [2][128]int{[128]int{}, [128]int{}}, 1+len(s)
	for i := range t {
		dp[1][t[i]]++
	}

	for i, j, count := 0, 0, 0; j < len(s); j++ {
		if dp[0][s[j]] < dp[1][s[j]] {
			count++
		}

		for dp[0][s[j]]++; i < j && dp[0][s[i]] > dp[1][s[i]]; {
			dp[0][s[i]], i = dp[0][s[i]]-1, i+1
		}

		if width := j - i + 1; count == len(t) && min > width {
			min, res = width, s[i:j+1]
		}
	}

	return
}

func minWindowMy(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	N := len(s)
	counter := map[byte]int{}
	for i := range t {
		counter[t[i]]++
	}
	win := map[byte]int{}
	check := func() bool {
		for k, v := range counter {
			if win[k] < v {
				return false
			}
		}
		return true
	}
	retLeft, retRight := -1, N+1
	for left, right := 0, 0; right < N; right++ {
		win[s[right]]++
		for ; left <= right && check(); left++ {
			if retRight-retLeft > right-left {
				retRight, retLeft = right, left
			}
			win[s[left]]--
		}
	}
	if retLeft == -1 {
		return ""
	}
	return s[retLeft : retRight+1]
}
func main() {

}
