// https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/

type Node struct {
    Val int
    Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var curentLevel []*Node
    curentLevel = append(curentLevel,root)
	var nextLevel []*Node
	var depth = 0
	for {
		if len(curentLevel) == 0{
			break
		}
		for _, node := range curentLevel{
			if len(node.Children) > 0 {
				nextLevel = append(nextLevel,node.Children...)
			}
		}
        depth = depth + 1
		curentLevel = nextLevel
		nextLevel = nil
	}
	return depth
}