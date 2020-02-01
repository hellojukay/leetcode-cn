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

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var levels [][]*TreeNode
	var current []*TreeNode
	var next []*TreeNode
	current = append(current, root)
	levels = append(levels, []*TreeNode{root})
	for {
		if len(current) == 0 {
			break
		}
		levels = append(levels, current)
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
	var result [][]int
	for i := len(levels) - 1; i > 0; i-- {
		level := levels[i]
		var vars []int
		for _, node := range level {
			vars = append(vars, node.Val)
		}
		result = append(result, vars)
	}
	return result
}
