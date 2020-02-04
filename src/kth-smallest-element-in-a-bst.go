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

//https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
func kthSmallest(root *TreeNode, k int) int {
	var arr = tree2arr(root)
	return arr[k-1]
}
func tree2arr(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	if root.Left != nil {
		result = append(tree2arr(root.Left), result...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, tree2arr(root.Right)...)
	}
	return result
}
