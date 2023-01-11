package main

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, array := range matrix {
		if inArray(array, target) {
			return true
		}
	}
	return false
}

func inArray(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if target < nums[0] {
		return false
	}
	if target > nums[len(nums)-1] {
		return false
	}
	var mid = len(nums) / 2
	if nums[mid] == target {
		return true
	}
	if mid+1 < len(nums) {
		if inArray(nums[mid+1:], target) {
			return true
		}
	}
	return inArray(nums[:mid], target)
}
