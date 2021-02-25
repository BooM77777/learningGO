package main

func canCross(stones []int) bool {
	if stones[1]-stones[0] != 1 {
		return false
	}
	deltaStep := []int{-1, 0, 1}
	stepHash := map[int][]bool{}
	for _, stone := range stones {
		stepHash[stone] = make([]bool, 1111)
	}
	stepHash[stones[1]][1] = true
	for i := 1; i < len(stones); i++ {
		var next int
		for step, ok := range stepHash[stones[i]] {
			if !ok {
				continue
			}
			for _, delta := range deltaStep {
				next = stones[i] + step + delta
				if _, ok := stepHash[next]; ok {
					stepHash[next][step+delta] = true
				}
			}
		}
	}
	for _, step := range stepHash[stones[len(stones)-1]] {
		if step {
			return true
		}
	}
	return false
}

func main() {

}
