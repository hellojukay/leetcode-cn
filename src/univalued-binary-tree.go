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

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var vec []int
	var current []*TreeNode
	var next []*TreeNode
	current = append(current, root)
	for {
		if len(current) == 0 {
			break
		}
		for _, node := range current {
			vec = append(vec, node.Val)
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
	var tmp = vec[0]
	for _, n := range vec {
		if tmp != n {
			return false
		}
	}
	return true
}
