func reverseOddLevels(root *TreeNode) *TreeNode {
	var list = toList(root)
	for index, slice := range list {
		if index%2 == 1 {
			reverse(slice)
		}
	}
	return root
}

func toList(root *TreeNode) [][]*TreeNode {
	if root == nil {
		return nil
	}
	var result [][]*TreeNode
	var current []*TreeNode = []*TreeNode{root}
	var next []*TreeNode
	for len(current) > 0 {
		result = append(result, current)
		for _, node := range current {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		current = next
		next = nil
	}
	return result
}

func reverse(list []*TreeNode) {
	if len(list) == 0 {
		return
	}
	var i = 0
	var j = len(list) - 1
	for i < j {
		tmp := list[i].Val
		list[i].Val = list[j].Val
		list[j].Val = tmp
		i++
		j--
	}
}