package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var currentLevel []*Node
	var nextLevel []*Node
	var result [][]int
	var tmp []int
	currentLevel = append(currentLevel, root)
	for {
		if len(currentLevel) == 0 {
			break
		}
		for _, node := range currentLevel {
			tmp = append(tmp, node.Val)
			if len(node.Children) > 0 {
				for _, chid := range node.Children {
					nextLevel = append(nextLevel, chid)
				}
			}
		}
		result = append(result, tmp)
		currentLevel = nextLevel
		nextLevel = nil
		tmp = nil
	}
	return result
}
