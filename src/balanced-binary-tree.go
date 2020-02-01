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

// https://leetcode-cn.com/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var m = deep(root.Left) - deep(root.Right)
	if m > 1 || m < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func deep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left = deep(root.Left) + 1
	var right = deep(root.Right) + 1
	if left > right {
		return left
	}
	return right
}
