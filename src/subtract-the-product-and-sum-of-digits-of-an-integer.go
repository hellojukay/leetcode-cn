package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
func subtractProductAndSum(n int) int {
	var arr []int
	for {
		if n == 0 {
			break
		}
		arr = append(arr, n%10)
		n = n / 10
	}
	var sum = 0
	var j = 1
	for _, x := range arr {
		sum = sum + x
		j = j * x
	}
	return j - sum
}
