package main

func wordsTyping(sentence []string, rows int, cols int) int {
	N := len(sentence)
	nextWord, sentCounter := make([]int, N), make([]int, N)
	var iter int
	for i := 0; i < N; i++ {
		iter = 0
		nextWord[i] = i
		for {
			iter += len(sentence[nextWord[i]]) + 1
			if iter > cols+1 {
				break
			}
			nextWord[i]++
			if nextWord[i] == N {
				nextWord[i] = 0
				sentCounter[i]++
			}
		}
	}
	ret, next := 0, 0
	for i := 0; i < rows; i++ {
		ret += sentCounter[next]
		next = nextWord[next]
	}
	return ret
}
func main() {

}
