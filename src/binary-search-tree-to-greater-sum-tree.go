package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var array = bstToArray(root)
	var hash = make(map[int]int)
	var count = 0
	for i := len(array) - 1; i >= 0; i-- {
		count += array[i]
		hash[array[i]] = count
	}
	inOrderLoopTree(root, hash)
	return root
}

func inOrderLoopTree(root *TreeNode, hash map[int]int) {
	if root == nil {
		return
	}
	root.Val = hash[root.Val]
	inOrderLoopTree(root.Left, hash)
	inOrderLoopTree(root.Right, hash)
}
func bstToArray(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	result = append(bstToArray(root.Left), root.Val)
	result = append(result, bstToArray(root.Right)...)
	return result
}
