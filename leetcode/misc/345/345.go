package main

func isVowels(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
func reverseVowels(s string) string {

	left, right := 0, len(s)-1
	bytesSeq := []byte(s)
	for left < right {
		for ; left < right && isVowels(bytesSeq[left]) == false; left++ {
		}
		for ; right > left && isVowels(bytesSeq[right]) == false; right-- {
		}
		if left < right {
			bytesSeq[left], bytesSeq[right] = bytesSeq[right], bytesSeq[left]
			left, right = left+1, right-1
		}
	}
	return string(bytesSeq)
}
func main() {

}
