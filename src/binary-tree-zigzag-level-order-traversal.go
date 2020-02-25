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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var current []*TreeNode
	var next []*TreeNode
	var flag = true
	current = append(current, root)
	for {
		if len(current) == 0 {
			break
		}
		var tmp []int
		for _, node := range current {
			if flag {
				tmp = append(tmp, node.Val)
			} else {
				tmp = append([]int{node.Val}, tmp...)
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		flag = !flag
		current = next
		next = nil
		result = append(result, tmp)
		tmp = nil
	}
	return result
}
