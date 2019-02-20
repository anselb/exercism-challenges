// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"unicode"
)

// Abbreviate returns an acronym with all words in the input string.
func Abbreviate(stringInput string) string {
	acronym := string(unicode.ToUpper(rune(stringInput[0])))

	for ind, char := range stringInput {
		charRune := rune(stringInput[ind])
		if ind+1 != len(stringInput) {
			charRune = rune(stringInput[ind+1])
		}

		if (char == ' ' || char == '-') && unicode.IsLetter(charRune) {
			acronym += string(unicode.ToUpper(charRune))
		}
	}

	return acronym
}
