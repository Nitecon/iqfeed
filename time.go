package iqfeed

import "time"

// TimeMsg represents a current timestamp from the network.
type TimeMsg struct {
	TimeStamp time.Time
}

// UnMarshall sends the data into the usable struct for consumption by the application.
func (tm *TimeMsg) UnMarshall(d []byte) {
	t, _ := time.Parse("20060102 150405", string(d))
	tm.TimeStamp = t
}
