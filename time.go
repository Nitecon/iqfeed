package iqfeed

import "time"

// TimeMsg represents a current timestamp from the network.
type TimeMsg struct {
	TimeStamp time.Time
}

// UnMarshall sends the data into the usable struct for consumption by the application.
func (tm *TimeMsg) UnMarshall(d []byte, loc *time.Location) {
	t, _ := time.ParseInLocation("20060102 15:04:05", string(d), loc)
	tm.TimeStamp = t
}
