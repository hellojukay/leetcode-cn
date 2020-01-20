package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/
func getNoZeroIntegers(n int) []int {
	var a, b int
	for i := 1; i < n; i++ {
		if noZero(i) && noZero(n-i) {
			a = i
			b = n - i
		}
	}
	return []int{a, b}
}
func noZero(n int) bool {
	if n < 10 {
		return true
	}
	if n%10 == 0 {
		return false
	}
	return noZero(n / 10)
}
