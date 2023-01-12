package main

func maximumCount(nums []int) int {
	var pos int
	var neg int
	var zero int
	for index := range nums {
		if nums[index] > 0 {
			pos++
			continue
		}
		if nums[index] == 0 {
			zero++
			continue
		}
		break
	}
	pos = len(nums) - zero - neg
	if pos > neg {
		return pos
	}
	return neg
}
