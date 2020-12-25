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
	m, n := len(g), len(s)
	for i, j := 0, 0; i < m && j < n; j++ {
		if g[i] <= s[j] {
			ret++
		} else {
			i++
		}

	}
	return ret
}
