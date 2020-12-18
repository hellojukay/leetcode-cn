package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var current []*Node
	var next []*Node
	current = append(current, root)
	for {
		if len(current) == 0 {
			break
		}
		for i := 0; i < len(current); i++ {
			if current[i].Left != nil {
				next = append(next, current[i].Left)
			}
			if current[i].Right != nil {
				next = append(next, current[i].Right)
			}
			if i == len(current)-1 {
				current[i].Next = nil
			} else {
				current[i].Next = current[i+1]
			}
		}
		current = next
		next = nil
	}
	return root
}
