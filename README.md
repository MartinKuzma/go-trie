# go-trie
Go-Trie is small library implementing trie datastructure with additional skipping optimization. It tries to build skipping indexes in nodes to avoid unnecessary checks. 
It is not suitable for use-cases with small dictionary (or single word searches). Naive implementations might be much better in those cases! **Always measure!**

# Installation
```go
import "github.com/MartinKuzma/go-trie"
```

# Features and decisions
- Case sensitive.
- Functions: `Contains`, `Find`
- Trie can be exported into JSON format and parsed from it, thus saving us from building and optimization steps.
- Nodes in trie are not saved in map due to performance reasons, however binary search might be performed when node has many children (>128).
- Designed to be immutable after it has been created. This ensures thread-safety and less complexity.

# Examples
```go
textToSearch :=  `Lorem ipsum ...`

// Create optimized trie
myTrie := trie.NewTrie().
    WithWords(substrings...).
    Optimize(true).
    Build()

myTrie.Find(textToSearch, func(result trie.SearchResult) {
    // Do something useful with result
    fmt.Printf("Found %s at position %d\n", result.Word, result.Position)
})

if myTrie.IsContained(textToSearch) {
    // Atleast one of the words in our trie is present.
}
```

# Benchmarks
These results were achieved by using database with 3000 words. Naive implementation uses `strings.Index` and `strings.Contains` functions. While the results represent huge performance advantage, we can easily find situations where naive implementation is much faster. 

Always use measurements that fit your data and use-case.

Trie creation and optimization included
```
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkNaiveFind-4         	      78	  13229589 ns/op	   81896 B/op	      12 allocs/op
BenchmarkTrieFind-4          	    2264	    518443 ns/op	   82286 B/op	      21 allocs/op
BenchmarkNaiveContained-4    	    4615	    248792 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrieIsContained-4   	 3948870	       299.8 ns/op	       0 B/op	       0 allocs/op
```


Trie creation and optimization not included
```
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkNaiveFind-4         	      78	  13386731 ns/op	   81896 B/op	      12 allocs/op
BenchmarkTrieFind-4          	    2245	    530640 ns/op	   81897 B/op	      12 allocs/op
BenchmarkNaiveContained-4    	    4566	    248930 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrieIsContained-4   	 4002381	       297.0 ns/op	       0 B/op	       0 allocs/op
```

# Future improvements
- Improve performance of optimization algorithm 
- Provide better examples and tests
- Implement naive version for small-sized database
  - Find out what constitues as a small database
