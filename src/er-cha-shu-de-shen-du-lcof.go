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

// https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left = 1 + maxDepth(root.Left)
	var rigth = 1 + maxDepth(root.Right)
	if left > rigth {
		return left
	}
	return rigth
}
