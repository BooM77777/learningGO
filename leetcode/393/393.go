package main

func validUtf8(data []int) bool {

	N := len(data)
	mask1, mask2, numberOfBytesToProcess := 1<<7, 1<<6, 0

	for i := 0; i < N; i++ {
		if numberOfBytesToProcess == 0 {

			// 检查 11100000
			mask := mask1
			for data[i]&mask != 0 {
				numberOfBytesToProcess++
				mask = mask >> 1
			}
			// 如果以0开头，直接过
			if numberOfBytesToProcess == 0 {
				continue
			}
			// 如果大于四个或者小于等于1个，直接gg
			if numberOfBytesToProcess > 4 || numberOfBytesToProcess == 1 {
				return false
			}
		} else {
			// 不是 10xxxxxx
			if (data[i]&mask1) == 0 || (data[i]&mask2) != 0 {
				return false
			}
		}
		numberOfBytesToProcess--
	}
	return numberOfBytesToProcess == 0
}

func main() {

}
