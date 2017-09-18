package trie

import "errors"

type Character string
type Keyword string

type Trie struct {
	s           map[Character]*Trie
	IndexedWord string
}

func NewTrie() *Trie {
	return &Trie{s: make(map[Character]*Trie), IndexedWord: ""}
}

func (t *Trie) IsKeyword() bool {
	return t.IndexedWord != ""
}

func (t *Trie) Retrieve(c Character) (*Trie, error) {
	if value, ok := t.s[c]; ok {
		return value, nil
	}
	return nil, errors.New("no item in trie")
}

func (t *Trie) Index(word Keyword) error {

	var currentTrie = t
	for _, char := range word {
		trie, err := currentTrie.Retrieve(Character(char))
		if err != nil {
			currentTrie.s[Character(char)] = NewTrie()
			currentTrie = currentTrie.s[Character(char)]
		} else {
			currentTrie = trie
		}

	}
	currentTrie.IndexedWord = string(word)
	return nil
}

func (t *Trie) IsCharIn(c Character) bool {
	_, ok := t.s[c]
	return ok
}
