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

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var current []*TreeNode
	var next []*TreeNode
	current = append(current, root)
	var list []int
	for {
		if len(current) == 0 {
			break
		}
		for _, node := range current {
			list = append(list, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		current = next
		next = nil
		result = append(result, list)
		list = nil
	}
	return result
}
