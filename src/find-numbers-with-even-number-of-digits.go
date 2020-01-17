package main

// https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits/
func findNumbers(nums []int) int {
	var result = 0
	for _, n := range nums {
		if isEven(n) {
			result++
		}
	}
	return result
}

func isEven(n int) bool {
	if n >= 0 && n < 10 {
		return false
	}
	if n < 0 && n > -10 {
		return false
	}
	return !isEven(n / 10)
}
