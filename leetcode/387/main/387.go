package main

import (
	"fmt"
)

type pair struct {
	firstBeFound int
	nums         int
}

func main() {
	s := "aab"
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {

	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	hash := make(map[rune]*pair)
	for i, c := range s {
		if p, ok := hash[c]; !ok {
			hash[c] = &pair{i, 1}
		} else {
			p.nums++
		}
	}

	toBeFound := -1
	for _, v := range hash {
		if v.nums == 1 {
			if toBeFound == -1 {
				toBeFound = v.firstBeFound
			} else {
				toBeFound = min(toBeFound, v.firstBeFound)
			}
		}
	}
	return toBeFound
}
