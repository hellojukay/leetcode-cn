func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var current []*TreeNode = []*TreeNode{root}
	var next []*TreeNode
	var flag = false
	for len(current) > 0 {
		var tmp []int
		for _, node := range current {
			if !flag {
				tmp = append(tmp, node.Val)
			} else {
				tmp = append([]int{node.Val}, tmp...)
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		current = next
		next = nil
		flag = !flag
		result = append(result, tmp)
	}
	return result
}