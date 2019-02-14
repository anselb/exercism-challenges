package clock

import (
	"fmt"
	"math"
)

// Clock : Clock is struct that is composed of the hours and
// minutes for that clock, in integers.
type Clock struct {
	hour, minute int
}

// New : New is a function that inputs the number of hours
// and minutes a clock stuct should have, and returns a new
// clock stuct with the correspnding time formatted in
// "military time".
func New(hour, minute int) Clock {
	if minute > 59 {
		hoursCarriedOver := math.Floor(float64(minute) / 59.0)
		hour += int(hoursCarriedOver)
		minute -= int(hoursCarriedOver) * 60
	}

	if hour > 23 {
		hour = hour % 24
	}

	if minute < 0 {
		for minute < 0 {
			hour--
			minute += 60
		}
	}

	if hour < 0 {
		for hour < 0 {
			hour += 24
		}
	}

	newClock := Clock{hour, minute}
	return newClock
}

// String : String is a clock method that that returns a
// string formatted version of the clock struct
func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

// Add : Add is a clock method that adds minutes to the
// clock and returns a new clock with the updated time
func (clock Clock) Add(minutes int) Clock {
	newMinutes := clock.minute + minutes
	updatedClock := New(clock.hour, newMinutes)
	return updatedClock
}

// Subtract : Subtract is a clock method that subtracts
// minutes to the clock and returns a new clock with the
// updated time
func (clock Clock) Subtract(minutes int) Clock {
	newMinutes := clock.minute - minutes
	updatedClock := New(clock.hour, newMinutes)
	return updatedClock
}
