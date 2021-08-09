
func include(root *TreeNode, p *TreeNode) bool {
	if root == p {
		return true
	}
	if root == nil {
		return false
	}
	return include(root.Left, p) || include(root.Right, p)
}

func isCommanAncestor(root, p, q *TreeNode) bool {
	return include(root, p) && include(root, q)
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if isCommanAncestor(root.Left, p, q) {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if isCommanAncestor(root.Right, p, q) {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
