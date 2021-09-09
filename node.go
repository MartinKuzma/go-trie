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

	return binarySearch(n, char, childrenCount)
}

func binarySearch(n *Node, key byte, childrenCount int) *Node {
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
func (n *Node) HasWord() bool {
	return n.Word != -1
}
