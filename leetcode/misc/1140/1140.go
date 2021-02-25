package main

func function(s string, l int) int {
	hash := map[int]bool{}
	tmp, mask := 0, 1
	for i := 0; i < l; i++ {
		tmp, mask = tmp*26+int(s[i]-'a'), mask*26
	}
	hash[tmp] = true

	for i := 1; i < len(s)-l+1; i++ {
		tmp = tmp*26 - int(s[i-1]-'a') + int(s[i+l-1]-'a')
		if hash[tmp] {
			return i
		}
		hash[tmp] = true
	}
	return -1
}

func longestDupSubstring(s string) string {
	left, right := 0, len(s)-1
	var mid, start int
	for left < right {
		mid = (left + right) / 2
		start = function(s, mid)
		if start == -1 {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return s[start : start+left]
}

func main() {

}
