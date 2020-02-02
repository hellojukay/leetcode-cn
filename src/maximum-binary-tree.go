package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	var arr = []int{3, 2, 1, 6, 0, 5}
	constructMaximumBinaryTree(arr)
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/maximum-binary-tree/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	var index = find_max(nums)
	var root = &TreeNode{
		Val: nums[index],
	}
	if len(nums[:index]) > 0 {
		root.Left = constructMaximumBinaryTree(nums[:index])
	}
	if index < len(nums)-1 && len(nums[index+1:]) > 0 {
		root.Right = constructMaximumBinaryTree(nums[index+1:])
	}
	return root
}

func find_max(arr []int) int {
	var max = arr[0]
	var index int = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			index = i
			max = arr[i]
		}
	}
	return index
}
