package flash

import (
	"testing"
)

func TestDummy(t *testing.T) {
	words := NewKeywords()
	words.Add("New York")
	words.Add("Hello")
	words.Add("Tokyo")
	foundKeywords := words.Extract("New York and Tokyo are Cities")
	if foundKeywords[0] != "New York" {
		t.Fail()
	}
	if foundKeywords[1] != "Tokyo" {
		t.Fail()
	}

}
