package main

import "sort"

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	var lenght = len(nums)
	return nums[lenght-1]*nums[lenght-2] - (nums[0] * nums[1])
}
