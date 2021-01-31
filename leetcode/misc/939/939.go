package main

func pointToHash(point []int) int {
	return 40001*point[0] + point[1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func minAreaRect(points [][]int) (ret int) {
	if len(points) < 4 {
		return 0
	}
	ret = 0
	hash := map[int]bool{}
	for _, point := range points {
		hash[pointToHash(point)] = true
	}
	var point1, point2 []int
	var s int
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i][0] == points[j][0] || points[i][1] == points[j][1] {
				continue
			}
			point1 = []int{points[i][0], points[j][1]}
			point2 = []int{points[j][0], points[i][1]}
			if hash[pointToHash(point1)] && hash[pointToHash(point2)] {
				s = abs((points[i][0] - points[j][0]) * (points[i][1] - points[j][1]))
				if ret == 0 {
					ret = s
				} else {
					ret = min(ret, s)
				}
			}
		}
	}
	return
}

func main() {

}
