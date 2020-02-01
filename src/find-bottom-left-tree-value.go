package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/find-bottom-left-tree-value/
func findBottomLeftValue(root *TreeNode) int {
	var result = root.Val
	var current []*TreeNode = []*TreeNode{root}
	var next []*TreeNode
	for {
		if len(current) == 0 {
			break
		}
		result = current[0].Val
		for _, node := range current {
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
	return result
}
