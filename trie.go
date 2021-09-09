package trie

import (
	"encoding/json"

	"github.com/juju/errors"
)

type Trie struct {
	words []string
	root  *Node
}

type SearchResult struct {
	Word     string
	Position int
}

// Used for marshalling and unmarshalling of Trie structure
type jsonTrie struct {
	Words []string
	Root  *Node
}

// Find searches for all occurences in Trie datastructure.
func (trie *Trie) Find(text string, f func(item SearchResult)) {
	length := len(text)

	for startPosition := 0; startPosition < length; {
		var lastVisited *Node = trie.root
		for indx := startPosition; indx < length; indx += 1 {
			if node := lastVisited.FindChild(text[indx]); node != nil {
				// We found word!
				if node.Word != -1 {
					f(SearchResult{
						Word:     trie.words[node.Word],
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
		var lastVisited *Node = trie.root
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

func (trie *Trie) HasPrefix(prefix string) bool {
	length := len(prefix)

	var lastVisited *Node = trie.root
	for indx := 0; indx < length; indx += 1 {
		if node := lastVisited.FindChild(prefix[indx]); node != nil {
			lastVisited = node
		} else {
			return false
		}

	}

	return true
}

func (trie *Trie) WordsWithPrefix(prefix string) []string {
	length := len(prefix)

	if length == 0 {
		return nil
	}

	var lastVisited *Node = trie.root
	for indx := 0; indx < length; indx += 1 {
		if node := lastVisited.FindChild(prefix[indx]); node != nil {
			lastVisited = node
		} else {
			return nil
		}
	}

	if lastVisited != nil {
		return trie.collectWords(lastVisited)
	}

	return nil
}

func (trie *Trie) collectWords(n *Node) []string {
	var results []string
	if n == nil {
		return results
	}

	var nodesToSearch []*Node
	nodesToSearch = append(nodesToSearch, n)
	nodesCount := 1
	for nodesCount > 0 {
		currentNode := nodesToSearch[nodesCount-1]
		if currentNode.HasWord() {
			results = append(results, trie.words[currentNode.Word])
		}

		nodesToSearch = nodesToSearch[:nodesCount-1]
		for i := len(currentNode.Children) - 1; i >= 0; i-- {
			nodesToSearch = append(nodesToSearch, currentNode.Children[i])
		}

		nodesCount = len(nodesToSearch)
	}

	return results
}

func (trie *Trie) Words() []string {
	return trie.words
}

func (trie *Trie) SortedWords() []string {
	return trie.collectWords(trie.root)
}

func (trie *Trie) ToJson() ([]byte, error) {
	return json.Marshal(jsonTrie{
		Words: trie.words,
		Root:  trie.root,
	})
}

func FromJson(data []byte) (*Trie, error) {
	var trie jsonTrie
	err := json.Unmarshal(data, &trie)
	if err != nil {
		return nil, errors.Annotate(err, "Error while parsing Trie datastructure from json")
	}

	return &Trie{
		words: trie.Words,
		root:  trie.Root,
	}, nil
}
