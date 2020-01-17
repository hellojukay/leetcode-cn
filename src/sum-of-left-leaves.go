package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/sum-of-left-leaves/
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil && (root.Right == nil || (root.Right.Left == nil && root.Right.Right == nil)) {
		return root.Left.Val
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil && (root.Right != nil || (root.Right.Left != nil && root.Right.Right != nil)) {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	if root.Right == nil {
		return sumOfLeftLeaves(root.Left)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
