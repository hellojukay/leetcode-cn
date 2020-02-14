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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var left = depth(root.Left)
	var right = depth(root.Right)
	if left-right > 1 || left-right < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left = 1 + depth(root.Left)
	var right = 1 + depth(root.Right)
	if left > right {
		return left
	}
	return right
}
