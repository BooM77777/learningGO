package main

import (
	"fmt"
	"strings"
)

func main() {

	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, s))
}

func wordPattern(pattern string, s string) bool {
	map1 := make(map[byte]string)
	map2 := make(map[string]byte)

	var subStrList = strings.Split(s, " ")

	if len(pattern) != len(subStrList) {
		return false
	}
	if len(pattern) == 0 {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		tmpStr, ok1 := map1[pattern[i]]
		tmpPattern, ok2 := map2[subStrList[i]]
		if !ok1 && !ok2 {
			map1[pattern[i]] = subStrList[i]
			map2[subStrList[i]] = pattern[i]
		} else if ok1 && ok2 {
			if tmpStr != subStrList[i] || tmpPattern != pattern[i] {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
