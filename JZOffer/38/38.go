package main

func permutation(s string) []string {
	ret := []string{}
	ans := []byte(s)
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(s) {
			ret = append(ret, string(ans))
		} else {
			hash := map[byte]bool{}
			for i := index; i < len(s); i++ {
				if hash[ans[i]] {
					continue
				}
				ans[index], ans[i] = ans[i], ans[index]
				hash[ans[index]] = true
				dfs(index + 1)
				ans[index], ans[i] = ans[i], ans[index]
			}
		}
	}
	dfs(0)
	return ret
}

func main() {

}
