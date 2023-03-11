package main

func search(nums []int, target int) int {
	return find(nums, 0, len(nums)-1, target)
}

func find(nums []int, start int, end int, target int) int {
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	mid := start + (end-start)/2
	if nums[mid] == target {
		return mid
	}
	if mid == start || mid == end {
		return -1
	}
	if target < nums[mid] {
		return find(nums, start, mid, target)
	}
	return find(nums, mid, end, target)
}
