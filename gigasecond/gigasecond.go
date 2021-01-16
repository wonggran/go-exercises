/*
Package gigasecond contains one function AddGigasecond.
*/
package gigasecond

import "time"

/*
Returns the time.Time after a gigasecond passes time.Time t.
*/
func AddGigasecond(t time.Time) time.Time {
	const giga = 1e9
	return t.Add(giga * time.Second)
}
