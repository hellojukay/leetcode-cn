/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	nodeArray := treeToArray(root)
	for i := 1; i < len(nodeArray); i++ {
		var currentLevel = nodeArray[i]
		for j := 0; j+1 < len(currentLevel); j++ {
			currentLevel[j].Next = currentLevel[j+1]
		}
	}
	return nodeArray[0][0]
}

func treeToArray(root *Node) [][]*Node {
	var result [][]*Node
	var currentLevel []*Node = []*Node{root}
	for len(currentLevel) > 0 {
		result = append(result, currentLevel)
		var next []*Node
		for _, node := range currentLevel {
			if node == nil {
				continue
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		currentLevel = next
	}
	return result
}