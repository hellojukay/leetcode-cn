/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var mid = len(nums) / 2
	var root = &TreeNode{
		Val: nums[mid],
	}
	root.Left = sortedArrayToBST(nums[:mid])
	if mid+1 < len(nums) {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}
	return root
}