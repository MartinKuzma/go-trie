# go-trie
Go-Trie is small library implementing trie datastructure with additional skipping optimization. It tries to build skipping indexes in nodes to avoid unnecessary checks. 
It is not suitable for use-cases with small dictionary (or single word searches). Naive implementations might be much better in those cases! **Always measure!**

### Features and implementation decisions
- Case sensitive.
- Functions: `Contains`, `Find`
- Trie can be exported into JSON format and parsed from it, thus saving us from building and optimization steps.
- Nodes in trie are not saved in map due to performance reasons, however binary search might be performed when node has many children (>128).
- It is designed to be immutable after it has been created. This might change in future.

```go
textToSearch :=  `Lorem ipsum ...`

// Create optimized trie
trie := contains.NewTrie().
    WithWords(substrings...).
    Optimize(true).
    Build()

trie.Find([]byte(textToSearch), func(result contains.SearchResult) {
    // Do something useful with result
})
```

### Future improvements
- Improve performance of optimization algorithm 
- Provide better examples
- Hide fields
