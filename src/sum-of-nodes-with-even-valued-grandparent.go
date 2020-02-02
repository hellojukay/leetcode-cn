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

// https://leetcode-cn.com/problems/sum-of-nodes-with-even-valued-grandparent/
func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var sum = 0
	if depth(root) > 2 {
		if root.Val%2 == 0 {
			if root.Left != nil && root.Left.Left != nil {
				sum = sum + root.Left.Left.Val
			}
			if root.Left != nil && root.Left.Right != nil {
				sum = sum + root.Left.Right.Val
			}
			if root.Right != nil && root.Right.Left != nil {
				sum = sum + root.Right.Left.Val
			}
			if root.Right != nil && root.Right.Right != nil {
				sum = sum + root.Right.Right.Val
			}
		}
		sum = sum + sumEvenGrandparent(root.Left) + sumEvenGrandparent(root.Right)
	}
	return sum
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left = 1 + depth(root.Left)
	var rigth = 1 + depth(root.Right)
	if left > rigth {
		return left
	}
	return rigth
}
