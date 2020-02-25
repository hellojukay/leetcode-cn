package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/sqrtx/
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	var start = 0
	var end = x
	for {
		var mid = (start + end) / 2
		var a = mid * mid
		var b = (mid + 1) * (mid + 1)
		if a <= x && b > x {
			return mid
		}
		if a > x {
			end = mid
			continue
		}
		if a < x {
			start = mid
		}
	}
}
