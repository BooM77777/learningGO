package main

func combinationSum3(k int, n int) [][]int {
	answer := [][]int{}
	var dfs func(pos, sum, cnt int, res []int)
	dfs = func(pos, sum, cnt int, res []int) {
		if sum == n && cnt == k {
			tmp := make([]int, len(res))
			copy(tmp, res)
			answer = append(answer, tmp)
			return
		} else if sum > n {
			return
		} else if pos == 10 {
			return
		} else if cnt > k {
			return
		}
		dfs(pos+1, sum, cnt, res)
		dfs(pos+1, sum+pos, cnt+1, append(res, pos))
	}
	dfs(1, 0, 0, []int{})
	return answer
}

func main() {

}
