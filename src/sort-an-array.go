package main

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	if len(nums) == 2 {
		if nums[0] < nums[1] {
			return []int{nums[0], nums[1]}
		} else {
			return []int{nums[1], nums[0]}
		}
	}
	var lt []int
	var gt []int
	for _, n := range nums[1:] {
		if n < nums[0] {
			lt = append(lt, n)
		} else {
			gt = append(gt, n)
		}
	}
	var result = append(sortArray(lt), nums[0])
	result = append(result, sortArray(gt)...)
	return result
}
