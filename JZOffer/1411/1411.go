package main

import "math"

func numOfWays(n int) int {

	base := int(math.Pow(10, 9)) + 7

	aba, abc := 6, 6
	for i := 2; i <= n; i++ {
		aba, abc = 3*aba+2*abc, 2*aba+2*abc
		aba, abc = aba%base, abc%base
	}
	return (aba + abc) % base
}

func main() {

}
