/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func cloneTree(root *Node) *Node {
	if root == nil {
		return root
	}
	newRoot := new(Node)
	newRoot.Val = root.Val
	var newChildren []*Node
	for _, node := range root.Children {
		newNode := cloneTree(node)
		newChildren = append(newChildren, newNode)
	}
	newRoot.Children = newChildren
	return newRoot
}