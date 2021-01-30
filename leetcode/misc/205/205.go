package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	hashS := map[byte]byte{}
	hashT := map[byte]byte{}
	if len(s) != len(t) {
		return false
	}
	N := len(s)
	var tmpS, tmpT byte
	var okS, okT bool
	for i := 0; i < N; i++ {
		tmpT, okS = hashS[s[i]]
		tmpS, okT = hashT[t[i]]
		if okS == false && okT == false {
			hashS[s[i]], hashT[t[i]] = t[i], s[i]
		} else if okS == true && okT == true {
			if tmpT != t[i] || tmpS != s[i] {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("abb", "egg"))
}
