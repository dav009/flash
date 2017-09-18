# Flash

Fast Keyword extraction using [Aho–Corasick algorithm](https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm)  and Tries.

Flash is a Golang reimplementation of [Flashtext](https://github.com/vi3k6i5/flashtext),

This is meant to be used when you have a large number of words that you want to:
 - extract from text
 - search and replace

Flash is meant as a replacement for Regex, which in such cases can be extremely slow.

As a reference using go-flash with 10K keywords in a 1000 sentence text, took 7.3ms,
while using regexes took 1minute 37s.


| Sentences | Keywords | String.Contains | Regex    | Go-Flash |
|-----------|----------|-----------------|----------|----------|
| 1000      | 10K      | 1.0035s         | 2.72ms   | 1min 37s


## Warning

This is a toy-project for me to get more familiar with Golang
Please be-aware of potential issues.

## Usage

```go

import "github.com/dav009/flash"

words := flash.NewKeywords()
words.Add("New York")
words.Add("Hello")
words.Add("Tokyo")
foundKeywords := words.Extract("New York and Tokyo are Cities")
fmt.Println(foundKeywords)
// [New York, Tokyo]
```

