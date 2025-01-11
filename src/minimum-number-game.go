func numberGame(nums []int) []int {
	// slice sort nums
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var arra []int
	for len(nums) > 0 {
		var a = nums[0]
		var b = nums[1]
		arra = append(arra, b)
		arra = append(arra, a)
		nums = nums[2:]
	}
	return arra
}