package finished

// Definition for a Node.
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	// 用递归就好，遇到child的时候先去遍历child
	if root == nil || (root.Next == nil && root.Child == nil) {
		return root
	}
	head := root
	if root.Child == nil {
		node := flatten(root.Next)
		root.Next = node
		node.Prev = root
	} else {
		next := flatten(root.Next)
		child := flatten(root.Child)
		root.Next = child
		child.Prev = root
		if next != nil {
			for child.Next != nil {
				child = child.Next
			}
			child.Next = next
			next.Prev = child
		}
		root.Child = nil
	}
	return head
}
