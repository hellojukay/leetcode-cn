package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if !(gt(root.Right, root.Val) && lt(root.Left, root.Val)) {
		return false
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}

func gt(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val <= val {
		return false
	}
	return gt(root.Left, val) && gt(root.Right, val)
}
func lt(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val >= val {
		return false
	}
	return lt(root.Left, val) && lt(root.Right, val)
}
