package twelve

import (
	"fmt"
)

// Verse returns a verse from 'The Twelve Days of Christmas'
// corresponding to the input day
func Verse(day int) string {
	whichDay := []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth",
	}
	giftOnDay := []string{
		"a Partridge in a Pear Tree.",
		"two Turtle Doves, and ",
		"three French Hens, ",
		"four Calling Birds, ",
		"five Gold Rings, ",
		"six Geese-a-Laying, ",
		"seven Swans-a-Swimming, ",
		"eight Maids-a-Milking, ",
		"nine Ladies Dancing, ",
		"ten Lords-a-Leaping, ",
		"eleven Pipers Piping, ",
		"twelve Drummers Drumming, ",
	}
	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", whichDay[day-1])

	for num := day - 1; num >= 0; num-- {
		verse += giftOnDay[num]
	}

	return verse
}

// Song returns the lyrics to 'The Twelve Days of Christmas'
func Song() string {
	fullSong := ""

	for day := 1; day <= 12; day++ {
		fullSong += Verse(day) + "\n"
	}

	return fullSong
}
