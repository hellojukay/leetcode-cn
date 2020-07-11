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

// https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	return root
}
