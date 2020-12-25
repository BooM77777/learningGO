package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 2}
	fmt.Println(findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ret := 0
	j := -1
	for i := 0; i < len(g); i++ {
		for {
			if j++; j >= len(s) {
				break
			}
			if s[j] >= g[i] {
				ret++
				break
			}
		}
	}
	return ret
}
