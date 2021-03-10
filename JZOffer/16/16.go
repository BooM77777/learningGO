package main

func myPow(x float64, n int) float64 {
	var pow func(x float64, n int) float64
	pow = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		} else if n%2 == 1 {
			return myPow(x, n-1) * x
		} else {
			tmp := myPow(x, n/2)
			return tmp * tmp
		}
	}
	if n >= 0 {
		return pow(x, n)
	}
	return 1 / pow(x, -n)
}

func main() {

}
