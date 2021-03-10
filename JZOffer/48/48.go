package main

func lengthOfLongestSubstring(s string) int {
	win := map[byte]int{}
	retVal := 0
	for left, right := 0, 0; right < len(s); right++ {
		win[s[right]]++
		for left < right && len(win) < right-left+1 {
			win[s[left]]--
			if win[s[left]] == 0 {
				delete(win, s[left])
			}
			left++
		}
		if retVal < right-left+1 {
			retVal = right - left + 1
		}
	}
	return retVal
}

func main() {

}
