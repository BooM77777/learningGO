package main

func generateParenthesis(n int) []string {

	cache := make([][]string, 9)
	for i := range cache {
		cache[i] = make([]string, 0, 1024)
	}
	cache[0] = append(cache[0], "")
	var recursion func(n int) []string
	recursion = func(n int) []string {
		if len(cache[n]) != 0 {
			return cache[n]
		}
		for i := 0; i <= n-1; i++ {
			for _, strLeft := range recursion(i) {
				for _, strRight := range recursion(n - 1 - i) {
					cache[n] = append(cache[n], "("+strLeft+")"+strRight)
				}
			}
		}
		return cache[n]
	}

	return recursion(n)
}

func main() {

}
