//https://leetcode-cn.com/problems/concatenation-of-array/
func getConcatenation(nums []int) []int {
	return append(nums,nums...)
}