package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var current []*TreeNode
	var next []*TreeNode
	var depth = 1
	current = append(current, root)
	for {
		if len(current) == 0 {
			break
		}
		for _, node := range current {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				return depth
			}
		}
		current = next
		next = nil
		depth++
	}
	return depth
}
