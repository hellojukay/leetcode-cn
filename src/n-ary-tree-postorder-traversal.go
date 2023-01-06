package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	if len(root.Children) == 0 {
		return []int{root.Val}
	}
	var result []int
	for _, node := range root.Children {
		result = append(result, postorder(node)...)
	}
	result = append(result, root.Val)
	return result
}
