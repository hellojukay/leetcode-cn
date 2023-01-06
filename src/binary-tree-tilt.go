package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return abs(sum(root.Left)-sum(root.Right)) + findTilt(root.Left) + findTilt(root.Right)
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sum(root.Left) + root.Val + sum(root.Right)
}
