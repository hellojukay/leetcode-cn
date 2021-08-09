
func isMonotonic(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	var flag = nums[len(nums)-1] >= nums[0]
	for i := 1; i <= len(nums)-1; i++ {
		var t = nums[i] >= nums[i-1]
		if t == flag {
			continue
		} else {
			return false
		}
	}
	return true
}
