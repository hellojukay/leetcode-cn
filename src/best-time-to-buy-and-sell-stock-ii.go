package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var profile = 0
	for i := 1; i < len(prices); i++ {
		var p = prices[i] - prices[i-1]
		if p > 0 {
			profile = profile + p
		}
	}
	return profile
}
