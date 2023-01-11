/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.cn/problems/diameter-of-binary-tree/
func diameterOfBinaryTree(root *TreeNode) int {
	var max = 0
	var current = []*TreeNode{root}
	var next []*TreeNode
	for len(current) > 0 {
		for _, node := range current {
			var tmp = depth(node.Left) + depth(node.Right)
			if tmp > max {
				max = tmp
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
			current = next
			next = nil
		}
	}
	return max
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	var left = depth(root.Left) + 1
	var right = depth(root.Right) + 1
	if left > right {
		return left
	}
	return right
}
