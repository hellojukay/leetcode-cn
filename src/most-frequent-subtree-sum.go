package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	var hash = make(map[int]int)
	count(root, hash)
	var result []int
	var max = 0
	for k, v := range hash {
		if v > max {
			result = []int{k}
			max = v
			continue
		}
		if v == max {
			result = append(result, k)
		}
	}
	return result
}

func count(root *TreeNode, hash map[int]int) {
	if root == nil {
		return
	}
	s := sum(root)
	v, ok := hash[s]
	if ok {
		hash[s] = v + 1
	} else {
		hash[s] = 1
	}
	if root.Left != nil {
		count(root.Left, hash)
	}
	if root.Right != nil {
		count(root.Right, hash)
	}
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sum(root.Left) + root.Val + sum(root.Right)
}
