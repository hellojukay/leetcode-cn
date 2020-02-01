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

// https://leetcode-cn.com/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var cuurent []*TreeNode
	var next []*TreeNode
	cuurent = append(cuurent, root)
	for {
		if len(cuurent) == 0 {
			break
		}
		result = append(result, cuurent[len(cuurent)-1].Val)
		for _, node := range cuurent {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		cuurent = next
		next = nil
	}
	return result
}
