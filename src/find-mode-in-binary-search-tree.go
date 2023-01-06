package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var count = make(map[int]int)
	countVal(root, count)
	var result []int
	var max = 0
	for k, v := range count {
		if v > max {
			max = v
			result = []int{k}
			continue
		}
		if v == max {
			result = append(result, k)
		}
	}
	return result
}

func countVal(root *TreeNode, count map[int]int) {
	if root == nil {
		return
	}
	v, ok := count[root.Val]
	if ok {
		count[root.Val] = v + 1
	} else {
		count[root.Val] = 1
	}
	countVal(root.Left, count)
	countVal(root.Right, count)
}
