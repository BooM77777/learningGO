package main

func myPow(x float64, n int) float64 {
	var pow func(x float64, n int) float64
	pow = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		ret := myPow(x, n/2)
		if n%2 == 1 {
			return x * ret * ret
		} else {
			return ret * ret
		}
	}
	if n > 0 {
		return pow(x, n)
	} else {
		return pow(1/x, -n)
	}
}

func main() {

}
