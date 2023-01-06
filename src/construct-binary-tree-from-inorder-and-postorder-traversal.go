package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [[左子树中序] 根 【右子树中序】]
// [[子左子树后序] 【右子树后序】 根]
// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	var root = &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	if len(inorder) == 1 {
		return root
	}
	var inLeft []int
	var inRight []int
	var postLeft []int
	var postRight []int
	var found = false
	for index := range inorder {
		if inorder[index] != root.Val {
			if !found {
				inLeft = append(inLeft, inorder[index])
			} else {
				inRight = append(inRight, inorder[index])
			}
		} else {
			found = true
		}
	}
	for index := range postorder {
		if index == len(postorder)-1 {
			break
		}
		if len(postLeft) < len(inLeft) {
			postLeft = append(postLeft, postorder[index])
		} else {
			postRight = append(postRight, postorder[index])
		}
	}
	root.Left = buildTree(inLeft, postLeft)
	root.Right = buildTree(inRight, postRight)
	return root
}
