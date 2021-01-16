package clock

import (
	"fmt"
	"math"
)

// Clock represents a 24-hour clock.
type Clock struct {
	hour   int
	minute int
}

/* Use Fprintf at two digits of integers. */
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

var numHoursPerDay int = 24
var numMinutesPerHour int = 60

/* Returns the hour given that it may be negative or positive (or zero). */
func alignedHour(hours int) int {
	circled := int(math.Mod(float64(hours), float64(numHoursPerDay)))

	if circled < 0 {
		return numHoursPerDay + circled
	}

	return circled
}

/* Returns the number of hours and minutes that this number of minutes is
equivelant to. */
func equivelantHoursAndMinutes(minutes int) (int, int) {
	/* A negative number of minutes means you go back in time meaning you will
	end up at a prior hour as well at a prior minute. */
	var circledHours int

	if minutes < 0 {
		mins := minutes

		for mins < 0 {
			mins += numMinutesPerHour
			circledHours--
		}
	} else {
		circledHours = int(minutes / numMinutesPerHour)
	}

	circledMinutes := int(math.Mod(float64(minutes), float64(numMinutesPerHour)))

	var alignedMins int

	if circledMinutes < 0 {
		alignedMins = numMinutesPerHour + circledMinutes
	} else {
		alignedMins = circledMinutes
	}

	return circledHours, alignedMins

}

// New returns a new 24-hour aligned clock given an arbitrary number of hours and minutes. */
func New(hour int, minute int) Clock {
	circledHours, alignedMins := equivelantHoursAndMinutes(minute)
	alignedHrs := alignedHour(hour + circledHours)

	return Clock{alignedHrs, alignedMins}
}

func (c Clock) additive(minutes int) Clock {
	/* The first call to equivelant converts the number of minutes to add into mixed
	hour-minute form leaving the number of hours final. */
	circledHours, alignedMins := equivelantHoursAndMinutes(minutes)
	/* The second call obtains mixed form to obtain the final number of minutes and hours
	from the existing clock's number of minutes and the prior call's mixed form minute's. */
	newCircledHours, newAlignedMins := equivelantHoursAndMinutes(c.minute + alignedMins)

	c.hour = alignedHour(c.hour + circledHours + newCircledHours)
	c.minute = newAlignedMins

	return c
}

// Add returns the clock aligned after some number of minutes passing.
func (c Clock) Add(minutes int) Clock {
	return c.additive(minutes)
}

// Subtract returns the clock aligned after winding it back some number of minutes.
func (c Clock) Subtract(minutes int) Clock {
	return c.additive(-minutes)
}
