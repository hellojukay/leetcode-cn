

func term(root *TreeNode, pre string) []string {
	if root == nil {
		return []string{pre}
	}
	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%s->%d", pre, root.Val)}
	}
	var result []string
	if root.Left != nil {
		result = append(result, term(root.Left, fmt.Sprintf("%s->%d", pre, root.Val))...)
	}
	if root.Right != nil {
		result = append(result, term(root.Right, fmt.Sprintf("%s->%d", pre, root.Val))...)
	}
	return result
}
func binaryTreePaths(root *TreeNode) []string {
	var pre = fmt.Sprintf("%d", root.Val)
	if root.Left == nil && root.Right == nil {
		return []string{pre}
	}
	var (
		left  []string
		right []string
	)
	if root.Left != nil {
		left = term(root.Left, pre)
	}
	if root.Right != nil {
		right = term(root.Right, pre)
	}
	return append(left, right...)
}
