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

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var arr1 = tree2arr(root1)
	var arr2 = tree2arr(root2)
	var result []int
	var i = 0
	var j = 0
	for {
		if i == len(arr1) && j == len(arr2) {
			break
		}
		if i < len(arr1) {
			if j < len(arr2) {
				if arr1[i] < arr2[j] {
					result = append(result, arr1[i])
					i++
				} else {
					result = append(result, arr2[j])
					j++
				}
			} else {
				result = append(result, arr1[i])
				i++
			}
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	return result
}

func tree2arr(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var result []int
	if root.Left != nil {
		result = append(tree2arr(root.Left), result...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, tree2arr(root.Right)...)
	}
	return result
}
