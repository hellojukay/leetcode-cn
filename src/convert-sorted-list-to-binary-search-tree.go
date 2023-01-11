/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	var array []int
	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}
	return sortArrayToBST(array)
}

func sortArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var mid = len(nums) / 2
	var root = &TreeNode{
		Val: nums[mid],
	}
	root.Left = sortArrayToBST(nums[:mid])
	if mid+1 < len(nums) {
		root.Right = sortArrayToBST(nums[mid+1:])
	}
	return root
}