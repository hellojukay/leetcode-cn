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

// https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if val > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode{
				Val: val,
			}
		} else {
			insertIntoBST(root.Right, val)
		}
	} else {
		if root.Left == nil {
			root.Left = &TreeNode{
				Val: val,
			}
		} else {
			insertIntoBST(root.Left, val)
		}
	}
	return root
}
