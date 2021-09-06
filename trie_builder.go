package trie

type TrieBuilder struct {
	words    []string
	optimize bool
}

func NewTrie() *TrieBuilder {
	return &TrieBuilder{
		optimize: true,
	}
}

// AddWord adds new word to list of words we will be searching for
func (tb *TrieBuilder) AddWord(word string) *TrieBuilder {
	if len(word) != 0 {
		tb.words = append(tb.words, word)
	}
	return tb
}

// AddWord adds new word to list of words we will be searching for
func (tb *TrieBuilder) WithWords(words ...string) *TrieBuilder {
	for _, word := range words {
		if len(word) != 0 {
			tb.words = append(tb.words, word)
		}
	}
	return tb
}

func (tb *TrieBuilder) Optimize(optimize bool) *TrieBuilder {
	tb.optimize = optimize
	return tb
}

// Build returns build search trie
func (tb *TrieBuilder) Build() *Trie {
	var trie *Trie = &Trie{
		root:  newNode(),
		words: tb.words,
	}

	//Build our trie
	for indx, word := range trie.words {
		insertWord(trie, word, indx)
	}

	if tb.optimize {
		tb.optimizeSkiping(trie)
	}

	return trie
}

func insertWord(trie *Trie, word string, indx int) {
	currentNode := trie.root
	for _, char := range []byte(word) {
		if node := currentNode.FindChild(char); node != nil {
			currentNode = node
		} else {
			// Missing node, create one and add it to children.
			currentNode = currentNode.AddChild(char)
		}
	}

	//Last visited node is end of our word
	currentNode.Word = indx
}

// Calculates skipping, might take long time
func (tb *TrieBuilder) optimizeSkiping(trie *Trie) {
	// Calculate for regular words.
	for _, word := range trie.words {
		calculateForWord(trie, word)
	}
}

func calculateForWord(trie *Trie, word string) {
	charBytes := []byte(word)
	//Small words are not interesting
	if len(charBytes) <= 1 {
		return
	}

	calculatedSkips := make([]int, len(charBytes))

	for j := len(charBytes); j >= 2; j -= 1 {
		charactersToSkip := 1
		for i := 1; i < j-1; i += 1 {
			sliceLength := j - i
			matchedBytes, matched := lookup(trie, charBytes[i:j])
			if matched || matchedBytes == sliceLength {
				// We found submatch along the way.
				// Skip by our current offset
				charactersToSkip = i
				break
			} else {
				// No equivalent of this path in our trie.
				// We can definetly skip first byte.
				charactersToSkip = i + 1
			}
		}

		calculatedSkips[j-1] = charactersToSkip
		// There is no point to look further down.
		if charactersToSkip == 1 {
			for indx := 0; indx < j; indx += 1 {
				calculatedSkips[indx] = 1
			}
			break
		}
	}

	applySkips(trie, charBytes, calculatedSkips)
}

func lookup(trie *Trie, part []byte) (int, bool) {
	matched := false
	currentNode := trie.root
	i := 0
	for ; i < len(part); i++ {
		if node := currentNode.FindChild(part[i]); node != nil {
			if node.Word != -1 {
				matched = true
			}
			currentNode = node
		} else {
			break
		}
	}

	return i, matched
}

func applySkips(trie *Trie, word []byte, skips []int) {
	currentNode := trie.root

	for i := 0; i < len(word); i++ {
		if node := currentNode.FindChild(word[i]); node != nil {
			node.Skip = skips[i]
			currentNode = node
		} else {
			break
		}
	}
}
