func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if target <= nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	var mid = len(nums) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return searchInsert(nums[:mid], target)
	}
	return mid + searchInsert(nums[mid:], target)
}