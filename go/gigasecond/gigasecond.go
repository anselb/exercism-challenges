// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds a gigasecond to an inputted time and returns it.
func AddGigasecond(timeInput time.Time) time.Time {
	const Gigasecond = 1000000000
	secondsToAdd := time.Second * time.Duration(Gigasecond)

	timeInput = timeInput.Add(secondsToAdd)
	return timeInput
}
