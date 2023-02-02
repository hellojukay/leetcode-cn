package main

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] != nums[1] {
		return nums[0]
	}
	return singleNonDuplicate(nums[2:])
}
