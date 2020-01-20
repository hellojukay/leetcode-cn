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

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var curent []*TreeNode
	var next []*TreeNode
	curent = append(curent, root)
	for {
		if len(curent) == 0 {
			break
		}
		var nums []int
		for _, node := range curent {
			nums = append(nums, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		result = append(result, nums)
		curent = next
		next = nil
	}
	return result
}
