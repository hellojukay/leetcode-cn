/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isAncestor(root *TreeNode, p *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == p {
		return true
	}
	if root.Left == p.Left && root.Right == p.Right && root.Val == p.Val {
		return true
	}
	return isAncestor(root.Left, p) || isAncestor(root.Right, p)
}
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if p == q {
		return p
	}
	if isAncestor(root.Left, p) && isAncestor(root.Left, q) {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if isAncestor(root.Right, p) && isAncestor(root.Right, q) {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

