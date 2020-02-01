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

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var results []float64
	var current []*TreeNode
	var next []*TreeNode
	current = append(current, root)
	for {
		if len(current) == 0 {
			break
		}
		length := len(current)
		var sum float64
		for i := 0; i < length; i++ {
			sum = sum + float64(current[i].Val)
			if current[i].Left != nil {
				next = append(next, current[i].Left)
			}
			if current[i].Right != nil {
				next = append(next, current[i].Right)
			}
		}
		current = next
		next = nil
		results = append(results, sum/float64(length))
		sum = 0.0
	}
	return results
}
