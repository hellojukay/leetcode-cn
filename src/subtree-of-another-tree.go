package main

type TreeNode struct {
	Val int

	Left *TreeNode

	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if equal(root, subRoot) {
		return true
	}
	if root == nil {
		return false
	}
	if isSubtree(root.Left, subRoot) {
		return true
	}
	if isSubtree(root.Right, subRoot) {
		return true
	}
	return false
}

func equal(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil && subRoot != nil {
		return false
	}
	if root != nil && subRoot == nil {
		return false
	}
	if root.Val != subRoot.Val {
		return false
	}
	return equal(root.Left, subRoot.Left) && equal(root.Right, subRoot.Right)
}
