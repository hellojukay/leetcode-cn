package main

func maxProfit(prices []int) int {
	var minPrice = max(prices)
	var profile = 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		if profile < price-minPrice {
			profile = price - minPrice
		}
	}
	return profile
}
func max(arr []int) int {
	var m = arr[0]
	for _, n := range arr {
		if n > m {
			m = n
		}
	}
	return m
}

func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
