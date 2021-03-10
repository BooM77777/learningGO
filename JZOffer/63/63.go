package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	retVal := 0
	minPrice := prices[0]
	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		} else {
			delta := price - minPrice
			if retVal < delta {
				retVal = delta
			}
		}
	}
	return retVal
}

func main() {

}
