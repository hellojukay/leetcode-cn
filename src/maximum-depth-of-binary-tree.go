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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dleft = 1 + maxDepth(root.Left)
	var dright = 1 + maxDepth(root.Right)
	if dleft > dright {
		return dleft
	}
	return dright
}
