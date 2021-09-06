package contains

import (
	"encoding/json"
)

type Trie struct {
	Words []string
	Root  *Node
}

type SearchResult struct {
	Word     string
	Position int
}

// Find searches for all occurences in Trie datastructure.
func Find(trie *Trie, srcBytes []byte, f func(item SearchResult)) {
	length := len(srcBytes)

	for startPosition := 0; startPosition < length; {
		var lastVisited *Node = trie.Root
		for indx := startPosition; indx < length; indx += 1 {
			if node := lastVisited.LookupChild(srcBytes[indx]); node != nil {
				if node.Word != -1 {
					f(SearchResult{
						Word:     trie.Words[node.Word],
						Position: startPosition,
					})
				}

				lastVisited = node
			} else {
				break
			}
		}
		startPosition += lastVisited.Skip
	}
}

// Contains searches for first occurence of any word in Trie datastructure.
func (trie *Trie) Contains(srcBytes []byte) bool {
	length := len(srcBytes)

	for startPosition := 0; startPosition < length; {
		var lastVisited *Node = trie.Root
		for indx := startPosition; indx < length; indx += 1 {
			if node := lastVisited.LookupChild(srcBytes[indx]); node != nil {
				if node.Word != -1 {
					return true
				}
				lastVisited = node
			} else {
				break
			}
		}
		startPosition += lastVisited.Skip
	}

	return false
}

func (trie *Trie) ToJson() ([]byte, error) {
	return json.Marshal(trie)
}

func FromJson(data []byte) (*Trie, error) {
	trie := &Trie{}
	err := json.Unmarshal(data, &Trie{})
	return trie, err
}
