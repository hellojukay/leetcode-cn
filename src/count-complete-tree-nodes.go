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

func countNodes(root *TreeNode) int {
	var count = 0
	if root == nil {
		return 0
	}
	var curent []*TreeNode
	var next []*TreeNode
	curent = append(curent, root)
	for {
		if len(curent) == 0 {
			break
		}
		for _, node := range curent {
			count = count + 1
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		curent = next
		next = nil
	}
	return count
}
