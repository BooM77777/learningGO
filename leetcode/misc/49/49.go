package main

func groupAnagrams(strs []string) [][]string {
	hashAString := func(str string) [26]int {
		ret := [26]int{}
		for _, b := range str {
			ret[b-'a']++
		}
		return ret
	}
	hash := map[[26]int]int{}
	ret := [][]string{}
	var strHash [26]int
	for _, str := range strs {
		strHash = hashAString(str)
		if index, ok := hash[strHash]; ok {
			ret[index] = append(ret[index], str)
		} else {
			ret = append(ret, []string{str})
			hash[strHash] = len(ret) - 1
		}
	}
	return ret
}
func main() {

}
