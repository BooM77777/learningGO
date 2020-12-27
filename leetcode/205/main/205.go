package main

import (
	"fmt"
)

func main() {
	s := "egg"
	b := "add"
	fmt.Println(isIsomorphic(s, b))
}

func isIsomorphic(s string, t string) bool {

	n := len(s)
	hashFromStoT := make(map[byte]byte)
	hashFromTtoS := make(map[byte]byte)

	ret := true
	for i := 0; i < n; i++ {
		expectT, flagSHash := hashFromStoT[s[i]]
		expectS, flagTHash := hashFromTtoS[t[i]]
		if flagSHash == false && flagTHash == false {
			hashFromStoT[s[i]] = t[i]
			hashFromTtoS[t[i]] = s[i]
		} else if flagSHash && flagTHash {
			if expectT != t[i] || expectS != s[i] {
				ret = false
				break
			}
		} else {
			ret = false
			break
		}
	}

	return ret
}
