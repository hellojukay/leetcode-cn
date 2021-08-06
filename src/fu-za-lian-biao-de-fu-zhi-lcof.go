package main 

func term(head *Node, hash map[Node]*Node) *Node {
	if head == nil {
		return nil
	}
	var node = new(Node)
	hash[*head] = node
	node.Val = head.Val
	node.Next = term(head.Next, hash)
	return node
}
func copyRandomList(head *Node) *Node {
	var hash = make(map[Node]*Node)
	newHead := term(head, hash)
	p := newHead
	for {
		if head == nil {
			break
		}
		if head.Random == nil {

			p.Random = nil
		} else {
			p.Random = hash[*(head.Random)]
		}
		head = head.Next
		p = p.Next
	}
	return newHead
}