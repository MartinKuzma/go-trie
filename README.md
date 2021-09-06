# go-trie
Go-Trie is small library implementing trie datastructure with additional skipping optimization. It tries to build skipping indexes in nodes to avoid unnecessary checks. 
It is not suitable for use-cases with small dictionary (or single word searches). Naive implementations might be much better in those cases! **Always measure!**

# Installation
```go
import "github.com/MartinKuzma/go-trie/v1"
```

# Features and decisions
- Case sensitive.
- Functions: `Contains`, `Find`
- Trie can be exported into JSON format and parsed from it, thus saving us from building and optimization steps.
- Nodes in trie are not saved in map due to performance reasons, however binary search might be performed when node has many children (>128).
- Designed to be immutable after it has been created. This ensures less complexity and thread-safety.

# Examples
```go
textToSearch :=  `Lorem ipsum ...`

// Create optimized trie
trie := NewTrie().
    WithWords(substrings...).
    Optimize(true).
    Build()

trie.Find(textToSearch, func(result contains.SearchResult) {
    // Do something useful with result
})

if trie.IsContained(textToSearch) {
    // Atleast one of the words in our trie is present.
}
```

# Future improvements
- Improve performance of optimization algorithm 
- Provide better examples and tests
- Hide fields
- Implement naive version for small-sized database
