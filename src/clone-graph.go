package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	var cache = make(map[int]*Node)
	return cloneNode(node, cache)
}

func cloneNode(node *Node, cache map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	v, ok := cache[node.Val]
	if ok {
		return v
	}
	newNode := new(Node)
	newNode.Val = node.Val
	cache[newNode.Val] = newNode
	var newNeighbors []*Node
	for _, neighbor := range node.Neighbors {
		newNeighbors = append(newNeighbors, cloneNode(neighbor, cache))
	}
	return newNode
}
