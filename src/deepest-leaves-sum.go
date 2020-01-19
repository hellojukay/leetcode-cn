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

// https://leetcode-cn.com/problems/deepest-leaves-sum/
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	var current []*TreeNode
	var next []*TreeNode
	current = append(current, root)
	var sum = 0
	for {
		if len(current) == 0 {
			break
		}
		sum = 0
		for _, node := range current {
			sum = sum + node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		current = next
		next = nil
	}
	return sum
}
