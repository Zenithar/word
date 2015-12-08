package word_test

import (
	"testing"

	"github.com/zenithar/word"
)

func TestGoodWord(t *testing.T) {
	var goodWords = []string{
		"",
		"facebook",
		"google",
		"zenithar",
		"linkedin",
	}

	for _, w := range goodWords {
		if !word.IsWord(w) {
			t.Errorf("%s must be a recognized word.", w)
		}
	}
}

func TestBadWord(t *testing.T) {
	var badWords = []string{
		"okajwwpvwfxigcb",
		"bejqryeinqtvask",
		"vhjnivbojjkvkcf",
	}

	for _, w := range badWords {
		if word.IsWord(w) {
			t.Errorf("%s must not be a recognized word.", w)
		}
	}
}
