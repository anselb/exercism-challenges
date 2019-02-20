package sublist

import (
	"reflect"
)

// Relation is a string of the relation between two lists.
type Relation string

// Sublist returns the relation of the first inputted list to the
// second inputted list.
// The relations are: equal, unequal, superlist, sublist.
func Sublist(listOne, listTwo []int) Relation {
	var bigList, smallList []int

	// Set outer and inner lists for loops below
	if len(listOne) > len(listTwo) {
		bigList = listOne
		smallList = listTwo
	} else if len(listOne) < len(listTwo) {
		bigList = listTwo
		smallList = listOne
	} else {
		// If list length is the same, see if they are equal.
		if reflect.DeepEqual(listOne, listTwo) {
			return "equal"
		}
		return "unequal"
	}

	bigListLen := len(bigList)
	smallListLen := len(smallList)

	// Go through all the possbile starting elements in the big list that can
	// be compared to the starting element of the small list
	for bigListInd := 0; bigListInd < bigListLen-smallListLen+1; bigListInd++ {
		// Hold the outer loop's current index.
		bigListIndHold := bigListInd

		// Check that the elements in the small list match the elements in the
		// big list while the list has at least one number and while the
		// the small list index is less than the length of the small list.
		for smallListInd := 0; smallListLen > 0 && smallList[smallListInd] == bigList[bigListInd] && smallListInd < smallListLen; smallListInd++ {
			bigListInd++
			// When we get to the end of the small list, there must be
			// a sublist or superlist, so see which one it is.
			if smallListInd == smallListLen-1 {
				if reflect.DeepEqual(bigList, listOne) {
					return "superlist"
				}
				return "sublist"
			}
		}

		// If a list is empty, it must either be a superlist or a
		// sublist of the other list.
		if smallListLen == 0 {
			if reflect.DeepEqual(bigList, listOne) {
				return "superlist"
			}
			return "sublist"
		}

		// If the match did not work, go back to the held index of the
		// outer loop.
		bigListInd = bigListIndHold
	}
	return "unequal"
}
