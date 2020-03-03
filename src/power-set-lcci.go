package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/power-set-lcci/
func subsets(nums []int) [][]int {
	var length = len(nums)
	if length == 0 {
		return nil
	}
	var max = pow2(length) - 1
	var result [][]int
	for i := 0; i <= max; i++ {
		// 检查每一位是否有值
		var n uint = uint(i)
		var current []int
		for j := length - 1; j >= 0; j-- {
			if n%2 == 1 {
				current = append([]int{nums[j]}, current...)
			}
			n = n >> 1
		}
		result = append(result, current)
	}
	return result
}
func pow2(n int) int {
	var result = 1
	for i := 0; i < n; i++ {
		result = result * 2
	}
	return result
}
