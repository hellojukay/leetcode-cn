package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/majority-element/
func majorityElement(nums []int) int {
	var top = -1
	var length = len(nums)
	for i := top + 1; i < length; i++ {
		if top < 0 {
			top = 0
			nums[top] = nums[i]
			continue
		}
		if nums[top] != nums[i] {
			top--
			continue
		}
		top++
		nums[top] = nums[i]
	}
	return nums[0]
}
