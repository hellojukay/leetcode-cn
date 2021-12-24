/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(arr []int) int {
	var m = arr[0]
	for _, n := range arr {
		if n > m {
			m = n
		}
	}
	return m
}
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result [][]int
	var current []*TreeNode = []*TreeNode{root}
	var next []*TreeNode
	for {
		if len(current) == 0 {
			break
		}
		var tmp []int
		for _, node := range current {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		current = next
		next = nil
		result = append(result, tmp)
	}
	var tmp []int
	for _, arr := range result {
		tmp = append(tmp, max(arr))
	}
	return tmp
}