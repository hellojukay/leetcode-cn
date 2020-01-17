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

// https://leetcode-cn.com/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return equal(root.Left, root.Right)
}

func equal(src *TreeNode, dest *TreeNode) bool {
	if src == dest {
		return true
	}
	if src == nil {
		return false
	}
	if dest == nil {
		return false
	}
	if src.Left == nil && src.Right == nil && dest.Left == nil && dest.Right == nil && src.Val == dest.Val {
		return true
	}
	if src.Left == nil && src.Right == nil && dest.Left == nil && dest.Right == nil && src.Val != dest.Val {
		return false
	}
	if src.Val == dest.Val && equal(src.Left, dest.Right) && equal(src.Right, dest.Left) {
		return true
	}
	return false
}
