package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for _, str := range strs {
			if i >= len(str) || str[i] != strs[0][i] {
				return str[:i]
			}
		}
	}
	return strs[0]
}
func main() {

}
