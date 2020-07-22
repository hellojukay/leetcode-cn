package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

//https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var result []int
	result = append(result, root.Val)
	if len(root.Children) > 0 {
		for _, child := range root.Children {
			result = append(result, preorder(child)...)
		}
	}
	return result
}
