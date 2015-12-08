// Package word is the package containing the word related functions
package word

import "strings"

var (
	vowels = map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	phones = map[string]bool{
		"sh":  true,
		"ch":  true,
		"ph":  true,
		"sz":  true,
		"cz":  true,
		"sch": true,
		"rz":  true,
		"dz":  true,
	}
)

// IsWord returns true when word seems to be a word string
// based on http://stackoverflow.com/questions/18717536/in-python-how-can-i-distinguish-between-a-human-readible-word-and-random-string
func IsWord(word string) bool {
	if len(strings.TrimSpace(word)) > 0 {
		consecutiveVowels := 0
		consecutiveConsonents := 0

		lword := []rune(strings.ToLower(word))
		for idx, letter := range lword {
			vowel, _ := vowels[letter]

			if idx >= 1 {
				prev := lword[idx-1]
				prevVowel, _ := vowels[prev]
				if !vowel && letter == 'y' && !prevVowel {
					vowel = true
				}

				if prevVowel != vowel {
					consecutiveConsonents = 0
					consecutiveVowels = 0
				}
			}

			if vowel {
				consecutiveVowels++
			} else {
				consecutiveConsonents++
			}

			if consecutiveConsonents >= 3 && consecutiveVowels >= 3 {
				return false
			}

			if consecutiveConsonents == 3 {
				subStr := lword[idx-2 : idx+1]
				if ok, _ := phones[string(subStr)]; ok {
					consecutiveConsonents--
					continue
				}
				return false
			}
		}
	}

	return true
}
