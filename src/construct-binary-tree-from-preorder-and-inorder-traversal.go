package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
func buildTree(pre []int, in []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	var root = &TreeNode{
		Val: pre[0],
	}
	if len(pre) == 1 {
		root.Left = nil
		root.Right = nil
		return root

	}
	var preLeft []int
	var preRight []int
	var inLeft []int
	var intRight []int
	var found = false
	for index := range in {
		if in[index] != root.Val {
			if !found {
				inLeft = append(inLeft, in[index])
			} else {
				intRight = append(intRight, in[index])
			}
		} else {
			found = true
		}
	}
	for i := 1; i < len(pre); i++ {
		if len(preLeft) < len(inLeft) {
			preLeft = append(preLeft, pre[i])
		} else {
			preRight = append(preRight, pre[i])
		}
	}
	root.Left = buildTree(preLeft, inLeft)
	root.Right = buildTree(preRight, intRight)
	return root
}
