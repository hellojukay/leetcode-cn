package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	var root *TreeNode
	for _, n := range preorder {
		root = appendToTree(root, n)
	}
	return root
}

func appendToTree(root *TreeNode, n int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: n,
		}
	}
	if n < root.Val {
		root.Left = appendToTree(root.Left, n)
	} else {
		root.Right = appendToTree(root.Right, n)
	}
	return root
}
