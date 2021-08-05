//https://leetcode-cn.com/problems/build-array-from-permutation/
func buildArray(nums []int) []int {
	var newlist = make([]int,len(nums))
	for i := 0 ; i < len(nums);i++ {
		newlist[i] = nums[nums[i]]
	}
	return newlist
}