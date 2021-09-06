package trie

import (
	"encoding/json"

	"github.com/juju/errors"
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
func (trie *Trie) Find(text string, f func(item SearchResult)) {
	length := len(text)

	for startPosition := 0; startPosition < length; {
		var lastVisited *Node = trie.Root
		for indx := startPosition; indx < length; indx += 1 {
			if node := lastVisited.FindChild(text[indx]); node != nil {
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

// IsContained searches for first occurence of any word in Trie datastructure.
func (trie *Trie) IsContained(text string) bool {
	length := len(text)

	for startPosition := 0; startPosition < length; {
		var lastVisited *Node = trie.Root
		for indx := startPosition; indx < length; indx += 1 {
			if node := lastVisited.FindChild(text[indx]); node != nil {
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
	return trie, errors.Annotate(err, "Error while parsing Trie datastrcture from json")
}
