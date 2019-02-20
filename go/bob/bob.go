// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey returns Bob's response to an inputted remark.
func Hey(remark string) string {
	const alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	remark = strings.TrimSpace(remark)

	silence := remark == ""
	yelling := remark == strings.ToUpper(remark) && strings.ContainsAny(remark, alpha)
	question := false
	if len(remark) > 0 {
		question = remark[len(remark)-1:] == "?"
	}

	if question && yelling {
		return "Calm down, I know what I'm doing!"
	}
	if yelling {
		return "Whoa, chill out!"
	}
	if question {
		return "Sure."
	}
	if silence {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
