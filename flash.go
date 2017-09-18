package flash

import "github.com/dav009/flash/trie"

func extractKeywords(t *trie.Trie, sentence string) []string {
	matches := make([]string, 0)
	currentTrie := t
	//sequence_end_pos := 0
	idx := 0

	sentenceLen := len(sentence)

	for idx < sentenceLen {

		char := string(sentence[idx])
		//fmt.Println(string(char))
		// it is a boundary char (i.e: space)
		if isWordBoundarie(trie.Character(char)) {

			idx2, longestSequenceFound := checkIfMatch(currentTrie, sentence, idx)

			idx = idx2
			if longestSequenceFound != "" {

				matches = append(matches, longestSequenceFound)

			}
			currentTrie = t

		} else if insideTrie, _ := currentTrie.Retrieve(trie.Character(char)); insideTrie != nil {

			// if it is indexed in the current trie

			currentTrie = insideTrie

		} else {

			// if it is not index in the currrent trie

			currentTrie = t

			idy := idx + 1

			for idy < sentenceLen {

				char := sentence[idy]
				if isWordBoundarie(trie.Character(char)) {
					break
				}
				idy += 1
			}
			idx = idy
		}
		if idx+1 >= sentenceLen {
			if currentTrie.IsKeyword() {
				matches = append(matches, currentTrie.IndexedWord)
			}
		}

		idx += 1
	}

	return matches

}

type Keywords struct {
	t *trie.Trie
}

func NewKeywords() Keywords {
	return Keywords{trie.NewTrie()}
}

func (x Keywords) Extract(sentence string) []string {
	return extractKeywords(x.t, sentence)
}

func (x Keywords) Add(w string) {
	x.t.Index(trie.Keyword(w))
}

func isWordBoundarie(c trie.Character) bool {
	return c == "" || c == " " || c == "\t" || c == "\n"
}

func checkIfMatch(t *trie.Trie, sentence string, idx int) (int, string) {
	char := sentence[idx]
	sequenceFound := ""
	longestSequenceFound := ""
	if t.IsKeyword() {
		sequenceFound = t.IndexedWord
		longestSequenceFound = t.IndexedWord
	}
	if t.IsCharIn(trie.Character(char)) {
		// look for longest sequence from here
		nextTrie, _ := t.Retrieve(trie.Character(char))
		seqFound, sequenceEndpos := searchLongest(sentence, idx, nextTrie)
		longestSequenceFound = seqFound

		if longestSequenceFound != "" && sequenceFound != longestSequenceFound {
			idx = sequenceEndpos

		}
	}
	return idx, longestSequenceFound
}

func searchLongest(sentence string, idx int, t *trie.Trie) (string, int) {
	longestSequenceFound := ""
	sequenceEndpos := -1
	sentenceLen := len(sentence)
	idx = idx + 1
	currentTrie := t
	for idx < sentenceLen {
		char := sentence[idx]
		if (isWordBoundarie(trie.Character(char))) && currentTrie.IsKeyword() {
			longestSequenceFound = currentTrie.IndexedWord
			sequenceEndpos = idx
		}
		if trie, _ := currentTrie.Retrieve(trie.Character(char)); trie != nil {
			// if it is indexed in the current trie

			currentTrie = trie
		} else {
			break
		}
		idx = idx + 1
	}

	if idx >= sentenceLen {

		if currentTrie.IsKeyword() {
			longestSequenceFound = currentTrie.IndexedWord
			sequenceEndpos = idx
		}
	}

	return longestSequenceFound, sequenceEndpos

}
