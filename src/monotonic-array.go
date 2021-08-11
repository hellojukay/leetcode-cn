package main

func isMonotonic(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	var max = 0
	var min = 0
	var pre = nums[0]
	for _, n := range nums {
		var result = compare(pre, n)
		if result > max {
			max = result
			pre = n
			continue
		}
		if result < min {
			min = result
		}
		pre = n
	}
	return max-min <= 1
}

func compare(a, b int) int {
	if a < b {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}
