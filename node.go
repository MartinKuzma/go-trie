package trie

type Node struct {
	Key      byte // Key/Character this node represents
	Word     int  // index of word
	Skip     int  // characters to skip
	Children []*Node
}

func newNode() *Node {
	return &Node{
		Word: -1,
		Skip: 1,
	}
}

func newKeyNode(key byte) *Node {
	return &Node{
		Word: -1,
		Skip: 1,
		Key:  key,
	}
}

func (n *Node) AddChild(char byte) *Node {
	child := newKeyNode(char)
	n.Children = append(n.Children, child)

	// Insert sort
	for i := len(n.Children) - 1; i >= 1; i -= 1 {
		current := n.Children[i]
		next := n.Children[i-1]

		if current.Key < next.Key {
			n.Children[i] = next
			n.Children[i-1] = current
		} else {
			break
		}
	}

	return child
}

func (n *Node) FindChild(char byte) *Node {
	childrenCount := len(n.Children)

	if childrenCount == 0 {
		return nil
	}

	// Using full binary search yields worse results
	return semiBinarySearch(n, char, childrenCount)
}

func linearSearch(n *Node, key byte, from int, to int) *Node {
	for ; from < to; from += 1 {
		node := n.Children[from]
		if node.Key == key {
			return node
		}
	}

	return nil
}

func semiBinarySearch(n *Node, key byte, childrenCount int) *Node {
	var start int = 0
	var end int = childrenCount
	for start < end {
		middle := (end + start) >> 1
		pivot := n.Children[middle].Key
		if key < pivot {
			end = middle
		} else if key > pivot {
			start = middle + 1
		} else {
			return n.Children[middle]
		}
	}

	return nil
}
