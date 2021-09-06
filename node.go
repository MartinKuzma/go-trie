package contains

import (
	"sort"
)

type Node struct {
	Key      byte // Key of this node
	Word     int  // index of word
	Skip     int  // characters to skip
	Children []*Node
}

const (
	sortThreshold = 64
)

func newNode() *Node {
	return &Node{
		Word: -1,
		Skip: 1,
	}
}

func NewKeyNode(key byte) *Node {
	return &Node{
		Word: -1,
		Skip: 1,
		Key:  key,
	}
}

func (n *Node) AddChild(char byte) *Node {
	child := NewKeyNode(char)
	n.Children = append(n.Children, child)

	if len(n.Children) > sortThreshold {
		sort.Slice(n.Children, func(i, j int) bool {
			return n.Children[i].Key < n.Children[j].Key
		})
	}

	return child
}

func (n *Node) LookupChild(char byte) *Node {
	childrenCount := len(n.Children)

	if childrenCount <= sortThreshold {
		// Regular linear search
		for i := 0; i < childrenCount; i += 1 {
			node := n.Children[i]
			if node.Key == char {
				return node
			}
		}
		return nil
	}

	// Use binary search
	resultIdx := sort.Search(childrenCount, func(i int) bool {
		return n.Children[i].Key >= char
	})

	if resultIdx < childrenCount && n.Children[resultIdx].Key == char {
		return n.Children[resultIdx]
	} else {
		return nil
	}

}
